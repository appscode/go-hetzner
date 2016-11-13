package hetzner

import "time"

type Product struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   []string `json:"description"`
	Traffic       string   `json:"traffic"`
	Dist          []string `json:"dist"`
	Arch          []int    `json:"arch"`
	Lang          []string `json:"lang"`
	Price         string   `json:"price"`
	PriceSetup    string   `json:"price_setup"`
	PriceVat      string   `json:"price_vat"`
	PriceSetupVat string   `json:"price_setup_vat"`
}

type AuthorizedKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}
type HostKey struct {
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}
type Transaction struct {
	ID            string    `json:"id"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	ServerNumber  *string   `json:"server_number"`
	ServerIP      *string   `json:"server_ip"`
	AuthorizedKey []struct {
		Key *AuthorizedKey `json:"key"`
	} `json:"authorized_key"`
	HostKey []struct {
		Key *HostKey `json:"key"`
	} `json:"host_key"`
	Comment *string `json:"comment"`
	Product struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description []string `json:"description"`
		Traffic     string   `json:"traffic"`
		Dist        string   `json:"dist"`
		Arch        string   `json:"arch"`
		Lang        string   `json:"lang"`
	} `json:"product"`
}

type CreateTransactionRequest struct {
	ProductID     string   `url:"product_id"`
	AuthorizedKey []string `url:"authorized_key,brackets"`
	Password      string   `url:"password,omitempty"`
	Dist          string   `url:"dist,omitempty"`
	Arch          int      `url:"arch,omitempty"`
	Lang          string   `url:"lang,omitempty"`
	Comment       string   `url:"comment,omitempty"`
	Test          bool     `url:"test"`
}
