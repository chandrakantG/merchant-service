package models

import "time"

type Merchant struct {
	ID        uint       `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Code      string     `json:"code,omitempty"`
	Status    uint8      `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CreateMerchantRequest struct {
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Status uint8  `json:"status,omitempty"`
}

type MerchantsWithPagination struct {
	Merchants  []Merchant `json:"merchants,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	NextPage     bool `json:"next_page"`
	PreviousPage bool `json:"previous_page"`
	Page         int  `json:"page"`
}

type GetMerchantRequest struct {
	Page int `json:"page,omitempty"`
}
