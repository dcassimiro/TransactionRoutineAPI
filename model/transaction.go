package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionRequest struct {
	AccountID        int             `json:"account_id" validate:"required"`
	OperationsTypeID int             `json:"operation_type_id" validate:"required"`
	Amount           decimal.Decimal `json:"amount" validate:"required"`
}

type Transaction struct {
	TransactionID    int             `json:"transaction_id" db:"transaction_ID"`
	AccountID        int             `json:"account_id" db:"account_ID"`
	OperationsTypeID int             `json:"operation_type_id" db:"operationType_ID"`
	Amount           decimal.Decimal `json:"amount" db:"amount"`
	EventDate        time.Time       `json:"eventDate" db:"eventDate"`
}
