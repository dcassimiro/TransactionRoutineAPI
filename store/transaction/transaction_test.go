package transaction_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store/transaction"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/trerr"
)

var defaultDate = time.Now()

func Test_storeImpl_Create(t *testing.T) {
	var amount float32 = 123.45
	cases := map[string]struct {
		ExpectedErr error

		InputTransaction model.TransactionRequest
		PrepareMock      func(mock sqlmock.Sqlmock)
	}{
		"should return success": {
			InputTransaction: model.TransactionRequest{AccountID: 1, OperationsTypeID: 2, Amount: amount},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO transactions (account_ID, operationType_ID, amount)
					VALUES (?, ?, ?);
				`).
					WithArgs(1, 2, amount).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"should return an error with the message: 'Unable to create a new transaction'": {
			ExpectedErr: trerr.New(http.StatusInternalServerError, "Unable to create a new transaction", nil),

			InputTransaction: model.TransactionRequest{},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
				INSERT INTO transactions (account_ID, operationType_ID, amount)
				VALUES (?, ?, ?);
			`).
					WithArgs().
					WillReturnError(errors.New("Unable to create a new transaction"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := transaction.NewStore(db, nil)
			ctx := context.Background()

			id, err := store.Create(ctx, cs.InputTransaction)

			if err == nil && id == "" {
				t.Error(id)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_storeImpl_ReadOne(t *testing.T) {
	cases := map[string]struct {
		ExpectedData *model.Transaction
		ExpectedErr  error

		InputTransactionID string
		PrepareMock        func(mock sqlmock.Sqlmock)
	}{
		"should return success": {
			ExpectedData: &model.Transaction{
				TransactionID:    1,
				AccountID:        1,
				OperationsTypeID: 4,
				Amount:           123.45,
				EventDate:        defaultDate,
			},

			InputTransactionID: "1",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT
						transaction_ID,
						account_ID,
						operationType_ID,
						amount,
						eventDate
					FROM transactions
					WHERE
						transaction_ID = ?;
				`).
					WithArgs("1").
					WillReturnRows(
						test.NewRows("transaction_ID", "account_ID", "operationType_ID", "amount", "eventDate").
							AddRow(1, 1, 4, 123.45, defaultDate),
					)
			},
		},
		"should return an error with the message:: 'I didn't find a transaction with this id'": {
			ExpectedErr: trerr.New(http.StatusNotFound, "I didn't find a transaction with this id", map[string]string{"transactionId": "1"}),

			InputTransactionID: "1",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT
						transaction_ID,
						account_ID,
						operationType_ID,
						amount,
						eventDate
					FROM transactions
					WHERE
						transaction_ID = ?;
			`).
					WithArgs("1").
					WillReturnError(errors.New("I didn't find a transaction with this id"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := transaction.NewStore(db, nil)
			ctx := context.Background()

			transaction, err := store.ReadOne(ctx, cs.InputTransactionID)

			if diff := cmp.Diff(transaction, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
