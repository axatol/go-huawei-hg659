package huaweihg659

import (
	"context"
	"encoding/json"
)

type DeviceInfo struct {
	DeviceName      string `json:"DeviceName"`      // e.g. "HG659"
	SerialNumber    string `json:"SerialNumber"`    // e.g. "J3N8W10123456789"
	ManufacturerOUI string `json:"ManufacturerOUI"` // e.g. "00E0FC"
	UpTime          int64  `json:"UpTime"`          // e.g. 3697479
	SoftwareVersion string `json:"SoftwareVersion"` // e.g. "V100R00123456789"
	HardwareVersion string `json:"HardwareVersion"` // e.g. "VER.B"
}

func (c *Client) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {
	raw, err := c.do(ctx, "/api/system/deviceinfo", nil)
	if err != nil {
		return nil, err
	}

	var data DeviceInfo
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

type WifiNetwork struct {
	Numbers       string `json:"Numbers"`       // e.g. "10"
	WifiEnable    bool   `json:"WifiEnable"`    // e.g. true
	WifiFrequency int8   `json:"WifiFrequency"` // e.g. 5
	WifiSSID      string `json:"WifiSsid"`      // e.g. "Blah"
}

func (c *Client) GetBroadcastedWifiNetworks(ctx context.Context) ([]WifiNetwork, error) {
	raw, err := c.do(ctx, "/api/system/wizard_wifi", nil)
	if err != nil {
		return nil, err
	}

	var data []WifiNetwork
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type DeviceCount struct {
	ActiveDeviceNumbers int64 `json:"ActiveDeviceNumbers"` // e.g. 15
	DatacardNumber      int64 `json:"DatacardNumber"`      // e.g. 0
	LANActiveNumber     int64 `json:"LanActiveNumber"`     // e.g. 5
	PhoneNumber         int64 `json:"PhoneNumber"`         // e.g. 2
	PrinterNumbers      int64 `json:"PrinterNumbers"`      // e.g. 0
	USBNumbers          int64 `json:"UsbNumbers"`          // e.g. 0
	UserNumber          int64 `json:"UserNumber"`          // e.g. 2
}

func (c *Client) GetConnectedDeviceCount(ctx context.Context) (*DeviceCount, error) {
	raw, err := c.do(ctx, "/api/system/device_count", nil)
	if err != nil {
		return nil, err
	}

	var data DeviceCount
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

type InternetDiagnosis struct {
	ConnectionStatus     string `json:"ConnectionStatus"`       // e.g. "Connected"
	DNSServers           string `json:"DNSServers"`             // e.g. ""
	DefaultGateway       string `json:"DefaultGateway"`         // e.g. ""
	DownMaxBitRate       string `json:"DownMaxBitRate"`         // e.g. "1000"
	DuplexMode           string `json:"DuplexMode"`             // e.g. "Full"
	ErrReason            string `json:"ErrReason"`              // e.g. "Success"
	ExternalIPAddress    string `json:"ExternalIPAddress"`      // e.g. ""
	HasInternetWan       bool   `json:"HasInternetWan"`         // e.g. true
	LinkStatus           string `json:"LinkStatus"`             // e.g. "Up"
	MACAddress           string `json:"MACAddress"`             // e.g. ""
	MaxBitRate           string `json:"MaxBitRate"`             // e.g. "1000"
	Status               string `json:"Status"`                 // e.g. "Connected"
	StatusCode           string `json:"StatusCode"`             // e.g. "Connected"
	UpMaxBitRate         string `json:"UpMaxBitRate"`           // e.g. "1000"
	Uptime               int64  `json:"Uptime"`                 // e.g. 350840
	WANAccessType        string `json:"WANAccessType"`          // e.g. "Ethernet"
	IPv4Enable           bool   `json:"X_IPv4Enable"`           // e.g. true
	IPv6Address          string `json:"X_IPv6Address"`          // e.g. ""
	IPv6AddressingType   string `json:"X_IPv6AddressingType"`   // e.g. "SLAAC"
	IPv6ConnectionStatus string `json:"X_IPv6ConnectionStatus"` // e.g. "Pending Disconnect"
	IPv6DNSServers       string `json:"X_IPv6DNSServers"`       // e.g. ""
	IPv6DefaultGateway   string `json:"X_IPv6DefaultGateway"`   // e.g. ""
	IPv6Enable           bool   `json:"X_IPv6Enable"`           // e.g. false
	IPv6PrefixLength     int64  `json:"X_IPv6PrefixLength"`     // e.g. 0
	IPv6PrefixList       string `json:"X_IPv6PrefixList"`       // e.g. ""
}

func (c *Client) GetInternetDiagnosis(ctx context.Context) (*InternetDiagnosis, error) {
	raw, err := c.do(ctx, "/api/system/diagnose_internet", nil)
	if err != nil {
		return nil, err
	}

	var data InternetDiagnosis
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
