package account_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store/account"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/trerr"
)

func Test_storeImpl_Create(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputAccount model.AccountRequest
		PrepareMock  func(mock sqlmock.Sqlmock)
	}{
		"should return success": {
			InputAccount: model.AccountRequest{DocumentNumber: "123"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO accounts (document_number)
					VALUES (?);
				`).
					WithArgs("123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		"should return an error with the message: 'Unable to create a new account'": {
			ExpectedErr: trerr.New(http.StatusInternalServerError, "Unable to create a new account", nil),

			InputAccount: model.AccountRequest{},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`
					INSERT INTO accounts (document_number)
					VALUES (?);
				`).
					WithArgs().
					WillReturnError(errors.New("Unable to create a new account"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := account.NewStore(db, nil)
			ctx := context.Background()

			id, err := store.Create(ctx, cs.InputAccount)

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
		ExpectedData *model.Account
		ExpectedErr  error

		InputAccountID string
		PrepareMock    func(mock sqlmock.Sqlmock)
	}{
		"should return success": {
			ExpectedData: &model.Account{
				AccountID:      1,
				DocumentNumber: "123",
			},

			InputAccountID: "1",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT
						account_ID,
						document_number
					FROM accounts 
					WHERE
						account_ID = ?;
				`).
					WithArgs("1").
					WillReturnRows(
						test.NewRows("account_ID", "document_number").
							AddRow(1, "123"),
					)
			},
		},
		"should return an error with the message: 'I didn't find an account with that id'": {
			ExpectedErr: trerr.New(http.StatusNotFound, "I didn't find an account with that id", map[string]string{"accountId": "default-id"}),

			InputAccountID: "default-id",
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`
					SELECT
						account_ID,
						document_number
					FROM accounts 
					WHERE
						account_ID = ?;
				`).
					WithArgs("readone-id").
					WillReturnError(errors.New("I didn't find an account with that id"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := account.NewStore(db, nil)
			ctx := context.Background()

			report, err := store.ReadOne(ctx, cs.InputAccountID)

			if diff := cmp.Diff(report, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
