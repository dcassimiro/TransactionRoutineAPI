package model

type AccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type Account struct {
	AccountID      int    `json:"account_id" db:"account_ID"`
	DocumentNumber string `json:"document_number" db:"document_number"`
}
