package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/pismo/TransactionRoutineAPI/store/account"
	"github.com/pismo/TransactionRoutineAPI/store/transaction"
)

// Container model for exporting instantiated repositories
type Container struct {
	Account     account.Store
	Transaction transaction.Store
}

// Options struct of options for creating an instance of the repositories
type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

// New creates a new instance of the repositories
func New(opts Options) *Container {
	return &Container{
		Account:     account.NewStore(opts.Writer, opts.Reader),
		Transaction: transaction.NewStore(opts.Writer, opts.Reader),
	}
}
