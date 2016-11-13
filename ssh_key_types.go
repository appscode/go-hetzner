package hetzner

type SSHKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
	Data        string `json:"data"`
}

type SSHKeyCreateRequest struct {
	Name string `url:"name"`
	Data string `url:"data"`
}

type SSHKeyUpdateRequest struct {
	Fingerprint string
	Name        string `url:"name"`
}
