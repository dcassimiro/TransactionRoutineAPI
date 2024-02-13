package transaction

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/pismo/TransactionRoutineAPI/app"
	"github.com/pismo/TransactionRoutineAPI/mocks"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/test"
	"github.com/pismo/TransactionRoutineAPI/validator"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

var defaultDate time.Time

func Test_handler_create(t *testing.T) {
	amount := decimal.NewFromFloat(123.45)
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData string

		InputBody   string
		PrepareMock func(mock *mocks.MockTransactionApp)
	}{
		"should return success": {
			ExpectedData: `{"data":{"transaction_id":1,"account_id":1,"operation_type_id":4,"amount":"123.45","eventDate":"0001-01-01T00:00:00Z"}}`,

			InputBody: `{"account_id": 1, "operation_type_id": 4, "amount": 123.45}`,
			PrepareMock: func(mockApp *mocks.MockTransactionApp) {
				mockApp.EXPECT().Create(gomock.Any(), model.TransactionRequest{
					AccountID:        1,
					OperationsTypeID: 4,
					Amount:           amount,
				}).
					Times(1).
					Return(&model.Transaction{
						TransactionID:    1,
						AccountID:        1,
						OperationsTypeID: 4,
						Amount:           amount,
						EventDate:        defaultDate,
					}, nil)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {

			// setup server mock
			req := httptest.NewRequest(http.MethodPost, "/v1/transactions", strings.NewReader(cs.InputBody))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = validator.New()

			c := e.NewContext(req, rec)

			// setup mock
			ctrl, _ := test.NewController(t)

			mockApp := mocks.NewMockTransactionApp(ctrl)
			cs.PrepareMock(mockApp)

			h := &handler{
				apps: &app.Container{Transaction: mockApp},
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
