package account_test

import (
	"net/http"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/google/go-cmp/cmp"
	"github.com/pismo/TransactionRoutineAPI/app/account"
	"github.com/pismo/TransactionRoutineAPI/mocks"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/trerr"
)

var (
	defaultError = trerr.New(http.StatusInternalServerError, "ocorreu um erro", nil)
)

func Test_appImpl_Create(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputAccount model.AccountRequest
		PrepareMock  func(mockStore *mocks.MockAccountStore)
	}{
		"should return success": {
			ExpectedData: &model.Account{
				AccountID:      1,
				DocumentNumber: "123",
			},

			InputAccount: model.AccountRequest{DocumentNumber: "123"},
			PrepareMock: func(mockStore *mocks.MockAccountStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.AccountRequest{DocumentNumber: "123"}).
					Times(1).
					Return("1", nil)

				mockStore.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(&model.Account{
						AccountID:      1,
						DocumentNumber: "123",
					}, nil)
			},
		},
		"should return error upon creation": {
			ExpectedErr: defaultError,

			InputAccount: model.AccountRequest{DocumentNumber: "123"},
			PrepareMock: func(mockStore *mocks.MockAccountStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.AccountRequest{DocumentNumber: "123"}).
					Times(1).
					Return("1", defaultError)
			},
		},
		"should return an error when reading": {
			ExpectedErr: defaultError,

			InputAccount: model.AccountRequest{DocumentNumber: "123"},
			PrepareMock: func(mockStore *mocks.MockAccountStore) {
				mockStore.EXPECT().Create(gomock.Any(), model.AccountRequest{DocumentNumber: "123"}).
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
			mockStore := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mockStore)

			app := account.NewApp(&store.Container{Account: mockStore})

			company, err := app.Create(ctx, cs.InputAccount)

			if diff := cmp.Diff(company, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_appImpl_ReadOne(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputAccountID string
		PrepareMock    func(mockStore *mocks.MockAccountStore)
	}{
		"should return success": {
			ExpectedData: &model.Account{
				AccountID:      1,
				DocumentNumber: "123",
			},

			InputAccountID: "1",
			PrepareMock: func(mockStore *mocks.MockAccountStore) {
				mockStore.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(&model.Account{
						AccountID:      1,
						DocumentNumber: "123",
					}, nil)
			},
		},
		"should return an error": {
			ExpectedErr: defaultError,

			InputAccountID: "1",
			PrepareMock: func(mockStore *mocks.MockAccountStore) {
				mockStore.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(nil, defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockStore := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mockStore)

			app := account.NewApp(&store.Container{Account: mockStore})
			feira, err := app.ReadOne(ctx, cs.InputAccountID)
			if diff := cmp.Diff(feira, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}

}
