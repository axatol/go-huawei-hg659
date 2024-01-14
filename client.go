package huaweihg659

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"crypto/sha256"

	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

type Client struct {
	Address *url.URL
	jar     *cookiejar.Jar
	client  *http.Client
}

func NewClient(address string) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %s", err)
	}

	addr, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %s", err)
	}

	client := Client{
		Address: addr,
		jar:     jar,
		client:  &http.Client{Jar: jar},
	}

	return &client, nil
}

func (c *Client) do(ctx context.Context, path string, body io.Reader) ([]byte, error) {
	method := http.MethodGet
	if body != nil {
		method = http.MethodPost
	}

	req, err := http.NewRequestWithContext(ctx, method, c.Address.String()+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %s", err)
	}

	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err)
	}

	raw, _ = bytes.CutPrefix(raw, []byte("while (1)\n    ;\n/*"))
	raw, _ = bytes.CutPrefix(raw, []byte("while(1); /*"))
	raw, _ = bytes.CutSuffix(raw, []byte("\n*/"))
	raw, _ = bytes.CutSuffix(raw, []byte("*/"))
	return raw, nil
}

func EncodePassword(username, password, csrfParam, csrfToken string) string {
	passwordSHA256Hex := sha256.Sum256([]byte(password))
	passwordSHA256 := []byte(hex.EncodeToString(passwordSHA256Hex[:]))
	passwordBase64 := base64.URLEncoding.EncodeToString(passwordSHA256)
	hash := sha256.New()
	hash.Write([]byte(username))
	hash.Write([]byte(passwordBase64))
	hash.Write([]byte(csrfParam))
	hash.Write([]byte(csrfToken))
	encoded := hex.EncodeToString(hash.Sum(nil))
	return encoded
}

type loginResponse struct {
	CsrfParam     string `json:"csrf_param"`    // e.g. "Ys39tguDiyozxIqwPjMKDBNwPRClkUf"
	CsrfToken     string `json:"csrf_token"`    // e.g. "TRq27Nr2EK76z0AhLh1Pv5L0lCMhbVA"
	ErrorCategory string `json:"errorCategory"` // e.g. "ok"
	Level         int    `json:"level"`         // e.g. 1
	IsWizard      bool   `json:"IsWizard"`      // e.g. true
	IsFirst       bool   `json:"IsFirst"`       // e.g. true
}

func (c *Client) Login(ctx context.Context, username, password string) error {
	raw, err := c.do(ctx, "/html/index.html", nil)
	if err != nil {
		return fmt.Errorf("failed to get login page: %s", err)
	}

	node, err := html.Parse(bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("failed to parse login page: %s", err)
	}

	var csrfToken, csrfParam string
	var crawl func(*html.Node)
	crawl = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key != "name" {
					continue
				}

				if attr.Val == "csrf_param" {
					for _, attr := range n.Attr {
						if attr.Key == "content" {
							csrfParam = attr.Val
							break
						}
					}
				}

				if attr.Val == "csrf_token" {
					for _, attr := range n.Attr {
						if attr.Key == "content" {
							csrfToken = attr.Val
							break
						}
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			crawl(child)
		}
	}

	crawl(node)
	if csrfToken == "" || csrfParam == "" {
		return fmt.Errorf("failed to find csrf token and param")
	}

	body := map[string]interface{}{
		"csrf": map[string]any{
			"csrf_param": csrfParam,
			"csrf_token": csrfToken,
		},
		"data": map[string]any{
			"UserName":       username,
			"Password":       EncodePassword(username, password, csrfParam, csrfToken),
			"isDestroyed":    false,
			"isDestroying":   false,
			"isInstance":     true,
			"isObserverable": true,
		},
	}

	raw, err = json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal login request: %s", err)
	}

	raw, err = c.do(ctx, "/api/system/user_login", bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("failed to login: %s", err)
	}

	var response loginResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		return fmt.Errorf("failed to unmarshal login response: %s - %s", err, string(raw))
	}

	if response.ErrorCategory != "ok" {
		return fmt.Errorf("login failed: %s", response.ErrorCategory)
	}

	return nil
}

func (c *Client) SessionID() string {
	if c.jar == nil {
		return ""
	}

	for _, cookie := range c.jar.Cookies(c.Address) {
		if cookie.Name == "SessionID_R3" {
			return cookie.Value
		}
	}

	return ""
}
