package model

type OperationsTypes struct {
	OperationsTypeID int    `json:"operation_type_id" db:"operationType_ID"`
	Description      string `json:"description" db:"description"`
}
