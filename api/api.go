package api

import (
	"github.com/labstack/echo/v4"
	v1 "github.com/pismo/TransactionRoutineAPI/api/v1"
	"github.com/pismo/TransactionRoutineAPI/app"
)

// Options struct of options for creating an instance of the routes
type Options struct {
	Group *echo.Group
	Apps  *app.Container
}

// Register api instance
func Register(opts Options) {
	v1.Register(opts.Group, opts.Apps)
}
