package transaction

import (
	"context"

	"github.com/pismo/TransactionRoutineAPI/logger"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store"
)

type App interface {
	Create(ctx context.Context, transaction model.TransactionRequest) (*model.Transaction, error)
}

func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores *store.Container
}

func (s *appImpl) Create(ctx context.Context, transaction model.TransactionRequest) (*model.Transaction, error) {
	transaction.Amount = amount(transaction.OperationsTypeID, transaction.Amount)

	id, err := s.stores.Transaction.Create(ctx, transaction)
	if err != nil {
		logger.ErrorContext(ctx, "app.transaction.Create.Create: ", err.Error())
		return nil, err
	}

	data, err := s.stores.Transaction.ReadOne(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.transaction.Create.ReadOne: ", err.Error())
		return nil, err
	}

	return data, nil
}
