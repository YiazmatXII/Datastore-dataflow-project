package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	. "app/webhook"
)

func (h *Handler) AddCreateCalc(e *echo.Echo) {
	e.POST("/", h.createCalc)
}

func (h *Handler) createCalc(c echo.Context) (err error) {
	dto := new(CreateInfoDto)
	if err = c.Bind(dto); err != nil {
		return err
	}
	if err = c.Validate(dto); err != nil {
		return err
	}

	elem := Webhook{
		Fields: dto.Fields,
		Op:     dto.Op,
	}
	err = h.StoreWebhook.Put(&elem)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, elem)
}