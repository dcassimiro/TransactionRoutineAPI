package transaction_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pismo/TransactionRoutineAPI/app/transaction"
	"github.com/pismo/TransactionRoutineAPI/mocks"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/trerr"
	"go.uber.org/mock/gomock"
)

var (
	defaultError = trerr.New(http.StatusInternalServerError, "an error has occurred", nil)
)

func Test_appImpl_Create(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Transaction

		InputTransaction model.TransactionRequest
		PrepareMock      func(mockStore *mocks.MockTransactionStore)
	}{
		"should return success": {
			ExpectedData: &model.Transaction{
				TransactionID:    1,
				AccountID:        1,
				OperationsTypeID: 2,
				Amount:           -12.34,
			},

			InputTransaction: model.TransactionRequest{AccountID: 1, OperationsTypeID: 2, Amount: -12.34},
			PrepareMock: func(mockStore *mocks.MockTransactionStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.TransactionRequest{AccountID: 1, OperationsTypeID: 2, Amount: 12.34}).
					Times(1).
					Return("1", nil)

				mockStore.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(&model.Transaction{
						TransactionID:    1,
						AccountID:        1,
						OperationsTypeID: 2,
						Amount:           -12.34,
					}, nil)
			},
		},
		"should return error upon creation": {
			ExpectedErr: defaultError,

			InputTransaction: model.TransactionRequest{AccountID: 1, OperationsTypeID: 2},
			PrepareMock: func(mockStore *mocks.MockTransactionStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.TransactionRequest{AccountID: 1, OperationsTypeID: 2}).
					Times(1).
					Return("1", defaultError)
			},
		},
		"should return an error when reading": {
			ExpectedErr: defaultError,

			InputTransaction: model.TransactionRequest{AccountID: 1, OperationsTypeID: 2},
			PrepareMock: func(mockStore *mocks.MockTransactionStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.TransactionRequest{AccountID: 1, OperationsTypeID: 2}).
					Times(1).
					Return("1", nil)

				mockStore.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(nil, defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockStore := mocks.NewMockTransactionStore(ctrl)

			cs.PrepareMock(mockStore)

			app := transaction.NewApp(&store.Container{Transaction: mockStore})

			company, err := app.Create(ctx, cs.InputTransaction)

			if diff := cmp.Diff(company, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
