package transaction

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pismo/TransactionRoutineAPI/app"
	"github.com/pismo/TransactionRoutineAPI/logger"
	"github.com/pismo/TransactionRoutineAPI/model"
)

func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.create)
}

type handler struct {
	apps *app.Container
}

func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.TransactionRequest
	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.transaction.create.Bind: ", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.transaction.create.Validate: ", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	data, err := h.apps.Transaction.Create(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Data: data,
	})
}
