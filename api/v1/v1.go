package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/pismo/TransactionRoutineAPI/api/v1/account"
	"github.com/pismo/TransactionRoutineAPI/api/v1/transaction"
	"github.com/pismo/TransactionRoutineAPI/app"
)

func Register(g *echo.Group, apps *app.Container) {
	v1 := g.Group("/v1")

	account.Register(v1.Group("/accounts"), apps)
	transaction.Register(v1.Group("/transactions"), apps)

}
