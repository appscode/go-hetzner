package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#Wake_on_LAN
type WOLService interface {
	Create(serverIP string) (*WOL, *http.Response, error)
	Get(serverIP string) (*WOL, *http.Response, error)
}

type WOLServiceImpl struct {
	client *Client
}

var _ WOLService = &WOLServiceImpl{}

func (s *WOLServiceImpl) Create(serverIP string) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%v", serverIP)

	type Data struct {
		WOL *WOL `json:"wol"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, nil, &data, true)
	return data.WOL, resp, err
}

func (s *WOLServiceImpl) Get(serverIP string) (*WOL, *http.Response, error) {
	path := fmt.Sprintf("/wol/%v", serverIP)

	type Data struct {
		WOL *WOL `json:"wol"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.WOL, resp, err
}
