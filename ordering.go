package hetzner

import (
	"fmt"
	"net/http"
)

// See: https://wiki.hetzner.de/index.php/Robot_Webservice/en#Server_ordering
type OrderingService interface {
	ListProducts() ([]*Product, *http.Response, error)
	GetProduct(id string) (*Product, *http.Response, error)

	ListTransactions() ([]*Transaction, *http.Response, error)
	CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error)
	GetTransaction(id string) (*Transaction, *http.Response, error)
}

type OrderingServiceImpl struct {
	client *Client
}

var _ OrderingService = &OrderingServiceImpl{}

func (s *OrderingServiceImpl) ListProducts() ([]*Product, *http.Response, error) {
	path := "/order/server/product"

	type Data struct {
		Product *Product `json:"product"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*Product, len(data))
	for i, d := range data {
		a[i] = d.Product
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) GetProduct(id string) (*Product, *http.Response, error) {
	path := fmt.Sprintf("/order/server/product/%v", id)

	type Data struct {
		Product *Product `json:"product"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Product, resp, err
}

func (s *OrderingServiceImpl) ListTransactions() ([]*Transaction, *http.Response, error) {
	path := "/order/server/transaction"

	type Data struct {
		Transaction *Transaction `json:"transaction"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*Transaction, len(data))
	for i, d := range data {
		a[i] = d.Transaction
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error) {
	path := "/order/server/transaction"

	type Data struct {
		Transaction *Transaction `json:"transaction"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Transaction, resp, err
}

func (s *OrderingServiceImpl) GetTransaction(id string) (*Transaction, *http.Response, error) {
	path := fmt.Sprintf("/order/server/transaction/%v", id)

	type Data struct {
		Transaction *Transaction `json:"transaction"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)
	return data.Transaction, resp, err
}
