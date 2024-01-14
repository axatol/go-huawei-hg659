package huaweihg659

import (
	"context"
	"encoding/json"
)

type LANStatistics struct {
	ReceiveBytes   int64  `json:"receivebytes"`   // e.g. 1115521851
	SendBytes      int64  `json:"sendbytes"`      // e.g. 532335628
	ReceivePacket  int64  `json:"receivepacket"`  // e.g. 7775540
	SendDiscard    int64  `json:"senddiscard"`    // e.g. 0
	SendError      int64  `json:"senderror"`      // e.g. 0
	ID             string `json:"ID"`             // e.g. "LAN1"
	ReceiveError   int64  `json:"receiveerror"`   // e.g. 0
	ReceiveDiscard int64  `json:"receivediscard"` // e.g. 0
	SendPacket     int64  `json:"sendpacket"`     // e.g. 30813230
}

func (c *Client) GetLANStatistics(ctx context.Context) ([]LANStatistics, error) {
	raw, err := c.do(ctx, "/api/ntwk/lan_info", nil)
	if err != nil {
		return nil, err
	}

	var data []LANStatistics
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type WifiStatistics struct {
	ReceiveBytes   int64  `json:"receivebytes"`   // e.g. 1432834374
	SendBytes      int64  `json:"sendbytes"`      // e.g. 3297610140
	WifiMode       string `json:"wifimode"`       // e.g. "5GHz"
	SendDiscard    int64  `json:"senddiscard"`    // e.g. 0
	SendError      int64  `json:"senderror"`      // e.g. 19046
	Name           string `json:"name"`           // e.g. "Blah"
	ReceivePacket  int64  `json:"receivepacket"`  // e.g. 160109751
	SendPacket     int64  `json:"sendpacket"`     // e.g. 403167409
	ReceiveDiscard int64  `json:"receivediscard"` // e.g. 0
	ReceiveError   int64  `json:"receiveerror"`   // e.g. 0
}

func (c *Client) GetWifiStatistics(ctx context.Context) ([]WifiStatistics, error) {
	raw, err := c.do(ctx, "/api/ntwk/lan_info", nil)
	if err != nil {
		return nil, err
	}

	var data []WifiStatistics
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type UMTSStatistics struct {
	BytesReceived   int64  `json:"BytesReceived"`   // e.g. 0
	UpstreamRate    string `json:"UpstreamRate"`    // e.g. "0.000000"
	PacketsSent     int64  `json:"PacketsSent"`     // e.g. 0
	BytesSent       int64  `json:"BytesSent"`       // e.g. 0
	PacketsReceived int64  `json:"PacketsReceived"` // e.g. 0
	HasDatacard     bool   `json:"HasDatacard"`     // e.g. false
	DownstreamRate  string `json:"DownstreamRate"`  // e.g. "0.000000"
}

func (c *Client) GetUMTSStatistics(ctx context.Context) (*UMTSStatistics, error) {
	raw, err := c.do(ctx, "/api/ntwk/umts_info_st", nil)
	if err != nil {
		return nil, err
	}

	var data UMTSStatistics
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
