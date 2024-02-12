package account

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
	Create(ctx context.Context, account model.AccountRequest) (string, error)
	ReadOne(ctx context.Context, accountID string) (*model.Account, error)
}

// NewStore creates a new instance of the transaction repository
func NewStore(writer, reader *sqlx.DB) Store {
	return &storeImpl{writer, reader}
}

type storeImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func (r *storeImpl) Create(ctx context.Context, account model.AccountRequest) (string, error) {
	result, err := r.writer.ExecContext(ctx, `
	INSERT INTO accounts (document_number)
	VALUES (?);
`, account.DocumentNumber)
	if err != nil {
		logger.ErrorContext(ctx, "store.account.Create: ", err.Error())
		return "", trerr.New(http.StatusInternalServerError, "Unable to create a new account", nil)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.ErrorContext(ctx, "store.account.Create: ", err.Error())

	}
	return strconv.FormatInt(id, 10), nil
}

func (r *storeImpl) ReadOne(ctx context.Context, accountID string) (*model.Account, error) {
	account := new(model.Account)
	err := r.writer.GetContext(ctx, account, `
	SELECT
		account_ID,
		document_number
	FROM accounts 
	WHERE
		account_ID = ?;
`, accountID)
	if err != nil {
		logger.ErrorContext(ctx, "store.account.ReadOne: ", err.Error())
		return nil, trerr.New(http.StatusNotFound, "I didn't find an account with that id", map[string]string{
			"accountId": accountID,
		})
	}

	return account, nil
}
