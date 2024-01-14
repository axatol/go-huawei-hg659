package huaweihg659

import (
	"context"
	"encoding/json"
)

type LANDevice struct {
	Active                bool   `json:"Active"`                // e.g. true
	QOSClassID            string `json:"QosclassID"`            // e.g. ""
	DeviceMaxDownloadRate int64  `json:"DeviceMaxDownLoadRate"` // e.g. 0
	HostName              string `json:"HostName"`              // e.g. "Foo_Wireless"
	Active46              bool   `json:"Active46"`              // e.g. true
	LeaseTime             int64  `json:"LeaseTime"`             // e.g. 14867
	ID                    string `json:"ID"`                    // e.g. "InternetGatewayDevice.LANDevice.1.Hosts.Host.1."
	IPv6Addrs             []any  `json:"Ipv6Addrs"`             // e.g. []
	ClassQueue            int64  `json:"ClassQueue"`            // e.g. -1
	Layer2Interface       string `json:"Layer2Interface"`       // e.g. "SSID1"
	ActualName            string `json:"ActualName"`            // e.g. ""
	IPAddress             string `json:"IPAddress"`             // e.g. "192.168.1.12"
	PolicerID             string `json:"PolicerID"`             // e.g. ""
	Domain                string `json:"domain"`                // e.g. "InternetGatewayDevice.LANDevice.1.WLANConfiguration.1"
	DeviceDownRateEnable  bool   `json:"DeviceDownRateEnable"`  // e.g. false
	MACAddress            string `json:"MACAddress"`            // e.g. "AA:BB:CC:DD:EE:FF"
	ParentControlEnable   bool   `json:"ParentControlEnable"`   // e.g. false
	MACFilterID           string `json:"MacFilterID"`           // e.g. ""
	AddressSource         string `json:"AddressSource"`         // e.g. "DHCP"
	V6Active              bool   `json:"V6Active"`              // e.g. false
	IconType              string `json:"IconType"`              // e.g. ""
	IPv6Address           string `json:"IPv6Address"`           // e.g. ""
	VendorClassID         string `json:"VendorClassID"`         // e.g. "android-dhcp-13"
}

func (c *Client) ListKnownLANDevices(ctx context.Context) ([]LANDevice, error) {
	raw, err := c.do(ctx, "/api/system/HostInfo", nil)
	if err != nil {
		return nil, err
	}

	var data []LANDevice
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type DHCPReservation struct {
	Addr       string `json:"Yiaddr"`     // e.g. "192.168.1.41"
	ID         string `json:"ID"`         // e.g. "InternetGatewayDevice.LANDevice.1.LANHostConfigManagement.DHCPStaticAddress.5."
	ActualName string `json:"ActualName"` // e.g. "bar_Ethernet"
	Chaddr     string `json:"Chaddr"`     // e.g. "AA:BB:CC:DD:EE:FF"
	Enable     bool   `json:"Enable"`     // e.g. true
}

func (c *Client) ListDHCPReservation(ctx context.Context) ([]DHCPReservation, error) {
	raw, err := c.do(ctx, "/api/ntwk/lan_ipaddressreserve", nil)
	if err != nil {
		return nil, err
	}

	var data []DHCPReservation
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}
