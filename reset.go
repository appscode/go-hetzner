package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#Reset
type ResetService interface {
	List() ([]*Reset, *http.Response, error)
	Get(serverIP string) (*Reset, *http.Response, error)
	Create(req *ResetCreateRequest) (*Reset, *http.Response, error)
}

type ResetServiceImpl struct {
	client *Client
}

var _ ResetService = &ResetServiceImpl{}

func (s *ResetServiceImpl) List() ([]*Reset, *http.Response, error) {
	path := "/reset"

	type Data struct {
		Reset *Reset `json:"reset"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*Reset, len(data))
	for i, d := range data {
		a[i] = d.Reset
	}
	return a, resp, err
}

func (s *ResetServiceImpl) Get(serverIP string) (*Reset, *http.Response, error) {
	path := fmt.Sprintf("/reset/%v", serverIP)

	type Data struct {
		Reset *Reset `json:"reset"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Reset, resp, err
}

func (s *ResetServiceImpl) Create(req *ResetCreateRequest) (*Reset, *http.Response, error) {
	path := fmt.Sprintf("/reset/%v", req.ServerIP)

	type Data struct {
		ResetResp *ResetCreateResponse `json:"reset"`
	}

	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)

	out := &Reset {
		data.ResetResp.ServerIP,
		data.ResetResp.ServerNumber,
		[]string{data.ResetResp.Type
	}
	return out, resp, err
}
