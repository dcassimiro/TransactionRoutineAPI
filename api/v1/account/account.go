package account

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
	g.GET("/:accountId", h.readOne)
}

type handler struct {
	apps *app.Container
}

func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.AccountRequest
	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.account.create.Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.account.create.Validate", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	data, err := h.apps.Account.Create(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Data: data,
	})
}

func (h *handler) readOne(c echo.Context) error {
	ctx := c.Request().Context()

	accountId := c.Param("accountId")
	if accountId == "" {
		logger.ErrorContext(ctx, "api.v1.account.readOne", "o campo 'accountId' é obrigatório")
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	data, err := h.apps.Account.ReadOne(ctx, accountId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: data,
	})
}
