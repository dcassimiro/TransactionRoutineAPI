package account

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/pismo/TransactionRoutineAPI/app"
	"github.com/pismo/TransactionRoutineAPI/mocks"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/trerr"
	"github.com/pismo/TransactionRoutineAPI/validator"
	"go.uber.org/mock/gomock"
)

var defaultError = trerr.New(http.StatusBadRequest, "Invalid Request", nil)

func Test_handler_create(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData string

		InputBody   string
		PrepareMock func(mock *mocks.MockAccountApp)
	}{
		"should return success": {
			ExpectedData: `{"data":{"account_id":1,"document_number":"123"}}`,

			InputBody: `{"document_number": "123"}`,
			PrepareMock: func(mockApp *mocks.MockAccountApp) {
				mockApp.EXPECT().Create(gomock.Any(), model.AccountRequest{
					DocumentNumber: "123",
				}).
					Times(1).
					Return(&model.Account{
						AccountID:      1,
						DocumentNumber: "123",
					}, nil)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			// setup server mock
			req := httptest.NewRequest(http.MethodPost, "/v1/accounts", strings.NewReader(cs.InputBody))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = validator.New()

			c := e.NewContext(req, rec)

			// setup mock
			ctrl, _ := test.NewController(t)

			mockApp := mocks.NewMockAccountApp(ctrl)
			cs.PrepareMock(mockApp)

			h := &handler{
				apps: &app.Container{Account: mockApp},
			}

			if diff := cmp.Diff(h.create(c), cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(strings.TrimSpace(rec.Body.String()), cs.ExpectedData); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_handler_readOne(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData string

		InputAcountID string
		PrepareMock   func(mock *mocks.MockAccountApp)
	}{
		"should return success": {
			ExpectedData: `{"data":{"account_id":1,"document_number":"123"}}`,

			InputAcountID: "1",
			PrepareMock: func(mockApp *mocks.MockAccountApp) {
				mockApp.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(&model.Account{
						AccountID:      1,
						DocumentNumber: "123",
					}, nil)
			},
		},
		"should return error: the 'accountId' field is mandatory": {
			ExpectedErr: defaultError,

			PrepareMock: func(mockApp *mocks.MockAccountApp) {},
		},
		"should return error": {
			ExpectedErr: defaultError,

			InputAcountID: "1",
			PrepareMock: func(mockApp *mocks.MockAccountApp) {
				mockApp.EXPECT().ReadOne(gomock.Any(), "1").
					Times(1).
					Return(nil, defaultError)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			// setup server mock
			req := httptest.NewRequest(http.MethodGet, "/v1/accounts", nil)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = validator.New()

			c := e.NewContext(req, rec)
			c.SetParamNames("accountId")
			c.SetParamValues(cs.InputAcountID)

			// setup mock
			ctrl, _ := test.NewController(t)

			mockApp := mocks.NewMockAccountApp(ctrl)
			cs.PrepareMock(mockApp)

			h := &handler{
				apps: &app.Container{Account: mockApp},
			}

			if diff := cmp.Diff(h.readOne(c), cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(strings.TrimSpace(rec.Body.String()), cs.ExpectedData); diff != "" {
				t.Error(diff)
			}
		})
	}
}
