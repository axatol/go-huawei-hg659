package huaweihg659

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	Address *url.URL
	jar     *cookiejar.Jar
	client  *http.Client
}

type ClientOption func(*Client)

func WithHTTPRoundTriper(rt http.RoundTripper) ClientOption {
	return func(c *Client) {
		c.client.Transport = rt
	}
}

func NewClient(address string, options ...ClientOption) (*Client, error) {
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

	for _, option := range options {
		option(&client)
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
