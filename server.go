package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#Server
type ServerService interface {
	ListServers() ([]*ServerSummary, *http.Response, error)
	GetServer(serverIP string) (*Server, *http.Response, error)
	UpdateServer(req *ServerUpdateRequest) (*Server, *http.Response, error)

	GetCancellation(serverIP string) (*Cancellation, *http.Response, error)
	CancelServer(req *CancelServerRequest) (*Cancellation, *http.Response, error)
	WithdrawCancellation(serverIP string) (*http.Response, error)
}

type ServerServiceImpl struct {
	client *Client
}

var _ ServerService = &ServerServiceImpl{}

func (s *ServerServiceImpl) ListServers() ([]*ServerSummary, *http.Response, error) {
	path := "/server"

	type Data struct {
		Server *ServerSummary `json:"server"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*ServerSummary, len(data))
	for i, d := range data {
		a[i] = d.Server
	}
	return a, resp, err
}

func (s *ServerServiceImpl) GetServer(serverIP string) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%v", serverIP)

	type Data struct {
		Server *Server `json:"server"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Server, resp, err
}

func (s *ServerServiceImpl) UpdateServer(req *ServerUpdateRequest) (*Server, *http.Response, error) {
	path := fmt.Sprintf("/server/%v", req.ServerIP)

	type Data struct {
		Server *Server `json:"server"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Server, resp, err
}

func (s *ServerServiceImpl) GetCancellation(serverIP string) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%v/cancellation", serverIP)

	type Data struct {
		Cancellation *Cancellation `json:"cancellation"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Cancellation, resp, err
}

func (s *ServerServiceImpl) CancelServer(req *CancelServerRequest) (*Cancellation, *http.Response, error) {
	path := fmt.Sprintf("/server/%v/cancellation", req.ServerIP)

	type Data struct {
		Cancellation *Cancellation `json:"cancellation"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Cancellation, resp, err
}

func (s *ServerServiceImpl) WithdrawCancellation(serverIP string) (*http.Response, error) {
	path := fmt.Sprintf("/server/%v/cancellation", serverIP)

	return s.client.Call(http.MethodDelete, path, nil, nil, true)
}
