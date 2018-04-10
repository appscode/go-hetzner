package hetzner

type Reset struct {
	ServerIP        string   `json:"server_ip"`
	ServerNumber    int      `json:"server_number"`
	Type            []string `json:"type"`
	OperatingStatus string   `json:"operating_status"`
}

type ResetCreateRequest struct {
	ServerIP string
	Type     string `url:"type"`
}

type ResetCreateResponse struct {
	ServerIP     string `json:"server_ip"`
	ServerNumber int    `json:"server_number"`
	Type         string `json:"type"`
}
