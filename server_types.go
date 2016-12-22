package hetzner

import . "github.com/appscode/go/encoding/json/types"

type ServerSummary struct {
	ServerIP     string `json:"server_ip"`
	ServerNumber int    `json:"server_number"`
	ServerName   string `json:"server_name"`
	Product      string `json:"product"`
	Dc           string `json:"dc"`
	Traffic      string `json:"traffic"`
	Flatrate     bool   `json:"flatrate"`
	Status       string `json:"status"`
	Throttled    bool   `json:"throttled"`
	Cancelled    bool   `json:"cancelled"`
	PaidUntil    string `json:"paid_until"`
}

type Server struct {
	ServerSummary
	IP     []string `json:"ip"`
	Subnet []struct {
		IP   string `json:"ip"`
		Mask string `json:"mask"`
	} `json:"subnet"`
	Reset   bool `json:"reset"`
	Rescue  bool `json:"rescue"`
	Vnc     bool `json:"vnc"`
	Windows bool `json:"windows"`
	Plesk   bool `json:"plesk"`
	Cpanel  bool `json:"cpanel"`
	Wol     bool `json:"wol"`
}

type ServerUpdateRequest struct {
	ServerIP   string
	ServerName string `url:"server_name"`
}

type Cancellation struct {
	ServerIP                 string        `json:"server_ip"`
	ServerNumber             int           `json:"server_number"`
	ServerName               string        `json:"server_name"`
	EarliestCancellationDate string        `json:"earliest_cancellation_date"`
	Cancelled                bool          `json:"cancelled"`
	CancellationDate         string        `json:"cancellation_date"`
	CancellationReason       ArrayOrString `json:"cancellation_reason"`
}

type CancelServerRequest struct {
	ServerIP           string
	CancellationDate   string `url:"cancellation_date"`
	CancellationReason string `url:"cancellation_reason,omitempty"`
}
