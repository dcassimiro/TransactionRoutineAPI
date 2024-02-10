package account

import (
	"context"

	"github.com/pismo/TransactionRoutineAPI/logger"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store"
)

type App interface {
	Create(ctx context.Context, account model.AccountRequest) (*model.Account, error)
	ReadOne(ctx context.Context, accountID string) (*model.Account, error)
}

func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores *store.Container
}

func (s *appImpl) Create(ctx context.Context, account model.AccountRequest) (*model.Account, error) {
	id, err := s.stores.Account.Create(ctx, account)
	if err != nil {
		logger.ErrorContext(ctx, "app.account.Create.Create", err.Error())
		return nil, err
	}

	data, err := s.stores.Account.ReadOne(ctx, id)
	if err != nil {
		logger.ErrorContext(ctx, "app.account.Create.ReadOne", err.Error())
		return nil, err
	}

	return data, nil
}

func (s *appImpl) ReadOne(ctx context.Context, accountID string) (*model.Account, error) {
	account, err := s.stores.Account.ReadOne(ctx, accountID)
	if err != nil {
		logger.ErrorContext(ctx, "app.account.ReadOne.ReadOne", err.Error())
		return nil, err
	}

	return account, nil
}
