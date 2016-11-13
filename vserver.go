package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#vServer
type VServerService interface {
	Command(req *VServerCommandRequest) (*http.Response, error)
}

type VServerServiceImpl struct {
	client *Client
}

var _ VServerService = &VServerServiceImpl{}

func (s *VServerServiceImpl) Command(req *VServerCommandRequest) (*http.Response, error) {
	path := fmt.Sprintf("/vserver/%v/command", req.ServerIP)
	return s.client.Call(http.MethodPost, path, req, nil, true)
}
