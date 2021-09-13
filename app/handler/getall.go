package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) AddGetAllCalc(e *echo.Echo) {
	e.GET("/:id", h.getAllCalc)
}

func (h *Handler) getAllCalc(c echo.Context) (err error) {
	list, err := h.StoreWebhook.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}