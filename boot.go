package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#Boot
type BootService interface {
	GetConfig(serverIP string) (*Boot, *http.Response, error)

	GetLinuxConfig(serverIP string) (*Linux, *http.Response, error)
	ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error)
	DeactivateLinux(serverIP string) (*http.Response, error)
}

type BootServiceImpl struct {
	client *Client
}

var _ BootService = &BootServiceImpl{}

func (s *BootServiceImpl) GetConfig(serverIP string) (*Boot, *http.Response, error) {
	path := fmt.Sprintf("/boot/%v", serverIP)

	type Data struct {
		Boot *Boot `json:"boot"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Boot, resp, err
}

func (s *BootServiceImpl) GetLinuxConfig(serverIP string) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%v/linux", serverIP)

	type Data struct {
		Linux *Linux `json:"linux"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Linux, resp, err
}

func (s *BootServiceImpl) ActivateLinux(req *ActivateLinuxRequest) (*Linux, *http.Response, error) {
	path := fmt.Sprintf("/boot/%v/linux", req.ServerIP)

	type Data struct {
		Linux *Linux `json:"linux"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Linux, resp, err
}

func (s *BootServiceImpl) DeactivateLinux(serverIP string) (*http.Response, error) {
	path := fmt.Sprintf("/boot/%v/linux", serverIP)

	return s.client.Call(http.MethodDelete, path, nil, nil, true)
}
