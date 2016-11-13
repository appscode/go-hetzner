package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#SSH_keys
type SSHKeyService interface {
	List() ([]*SSHKey, *http.Response, error)
	Create(req *SSHKeyCreateRequest) (*SSHKey, *http.Response, error)
	Get(fingerprint string) (*SSHKey, *http.Response, error)
	Update(req *SSHKeyUpdateRequest) (*SSHKey, *http.Response, error)
	Delete(fingerprint string) (*http.Response, error)
}

type SSHKeyServiceImpl struct {
	client *Client
}

var _ SSHKeyService = &SSHKeyServiceImpl{}

func (s *SSHKeyServiceImpl) List() ([]*SSHKey, *http.Response, error) {
	path := "/key"

	type Data struct {
		Key *SSHKey `json:"key"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*SSHKey, len(data))
	for i, d := range data {
		a[i] = d.Key
	}
	return a, resp, err
}

func (s *SSHKeyServiceImpl) Create(req *SSHKeyCreateRequest) (*SSHKey, *http.Response, error) {
	path := "/key"

	type Data struct {
		Key *SSHKey `json:"key"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Key, resp, err
}

func (s *SSHKeyServiceImpl) Get(fingerprint string) (*SSHKey, *http.Response, error) {
	path := fmt.Sprintf("/key/%v", fingerprint)

	type Data struct {
		Key *SSHKey `json:"key"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Key, resp, err
}

func (s *SSHKeyServiceImpl) Update(req *SSHKeyUpdateRequest) (*SSHKey, *http.Response, error) {
	path := fmt.Sprintf("/key/%v", req.Fingerprint)

	type Data struct {
		Key *SSHKey `json:"key"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Key, resp, err
}

func (s *SSHKeyServiceImpl) Delete(fingerprint string) (*http.Response, error) {
	path := fmt.Sprintf("/key/%v", fingerprint)
	return s.client.Call(http.MethodDelete, path, nil, nil, true)
}
