package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// GetDB returns an instance of the mock db
func GetDB() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual), sqlmock.MonitorPingsOption(true))
	return sqlx.NewDb(db, "mysql"), mock
}

// NewRows returns a model for adding rows
func NewRows(columns ...string) *sqlmock.Rows {
	return sqlmock.NewRows(columns)
}
