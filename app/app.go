package app

import (
	"github.com/pismo/TransactionRoutineAPI/app/account"
	"github.com/pismo/TransactionRoutineAPI/app/transaction"
	"github.com/pismo/TransactionRoutineAPI/store"
)

type Container struct {
	Account     account.App
	Transaction transaction.App
}

type Options struct {
	Stores *store.Container
}

func New(opts Options) *Container {
	container := &Container{
		Account:     account.NewApp(opts.Stores),
		Transaction: transaction.NewApp(opts.Stores),
	}
	return container
}
