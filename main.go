package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/pismo/TransactionRoutineAPI/api"
	"github.com/pismo/TransactionRoutineAPI/app"
	"github.com/pismo/TransactionRoutineAPI/db"
	"github.com/pismo/TransactionRoutineAPI/logger"
	"github.com/pismo/TransactionRoutineAPI/model"
	"github.com/pismo/TransactionRoutineAPI/store"
	"github.com/pismo/TransactionRoutineAPI/validator"
)

const dbParameter = "transaction_db?charset=utf8mb4,utf8\\u0026readTimeout=30s\\u0026writeTimeout=30s&parseTime=true"

func main() {
	ec := echo.New()
	ec.Validator = validator.New()

	db.CreateDB()
	url := model.Url()

	dbWriter := sqlx.MustConnect("mysql", url+dbParameter)
	dbReader := sqlx.MustConnect("mysql", url+dbParameter)

	// creation of stores with the injection of the writing and reading database
	stores := store.New(store.Options{
		Writer: dbWriter,
		Reader: dbReader,
	})

	// creation of services
	apps := app.New(app.Options{
		Stores: stores,
	})

	// handler records
	api.Register(api.Options{
		Group: ec.Group(""),
		Apps:  apps,
	})

	ec.Start(":8080")
	logger.Info("API initialized successfully!!!")

	dbReader.Close()
	dbWriter.Close()
	ec.Close()
}
