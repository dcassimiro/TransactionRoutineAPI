package transaction

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pismo/TransactionRoutineAPI/logger"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/trerr"
)

type Store interface {
	Create(ctx context.Context, transaction model.TransactionRequest) (string, error)
	ReadOne(ctx context.Context, transactionID string) (*model.Transaction, error)
}

// NewStore creates a new instance of the transaction repository
func NewStore(writer, reader *sqlx.DB) Store {
	return &storeImpl{writer, reader}
}

type storeImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func (r *storeImpl) Create(ctx context.Context, transaction model.TransactionRequest) (string, error) {
	result, err := r.writer.ExecContext(ctx, `
	INSERT INTO transactions (account_ID, operationType_ID, amount)
	VALUES (?, ?, ?);
`, transaction.AccountID, transaction.OperationsTypeID, transaction.Amount)
	if err != nil {
		return "", trerr.New(http.StatusInternalServerError, "Unable to create a new transaction", nil)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.ErrorContext(ctx, "store.transaction.Create: ", err.Error())

	}
	return strconv.FormatInt(id, 10), nil
}

func (r *storeImpl) ReadOne(ctx context.Context, transactionID string) (*model.Transaction, error) {
	transaction := new(model.Transaction)
	err := r.writer.GetContext(ctx, transaction, `
	SELECT
		transaction_ID,
		account_ID,
		operationType_ID,
		amount,
		eventDate
	FROM transactions 
	WHERE
		transaction_ID = ?;
`, transactionID)
	if err != nil {
		logger.ErrorContext(ctx, "store.transaction.ReadOne: ", err.Error())
		return nil, err
	}

	return transaction, nil
}
