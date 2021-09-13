package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	. "app/webhook"
)

func (h *Handler) AddGetCalc(e *echo.Echo) {
	e.GET("/:id", h.getCalc)
}

func (h *Handler) getCalc(c echo.Context) (err error) {
	dto := new(GetInfoDto)
	if err = c.Bind(dto); err != nil {
		return err
	}
	if err = c.Validate(dto); err != nil {
		return err
	}

	elem := new(Webhook)

	elem, err = h.StoreWebhook.Get(dto.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, elem)
}