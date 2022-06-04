package models

import "time"

type MerchantTeam struct {
	ID         uint       `json:"id,omitempty"`
	MerchantID uint       `json:"merchant_id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Email      string     `json:"email,omitempty"`
	Status     uint8      `json:"status,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type CreateMerchantTeamRequest struct {
	MerchantID uint   `json:"merchant_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}

type MerchantTeamsWithPagination struct {
	MerchantTeams []MerchantTeam `json:"merchant_teams,omitempty"`
	Pagination    Pagination     `json:"pagination,omitempty"`
}
