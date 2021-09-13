package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	. "app/webhook"
)

func (h *Handler) AddUpdateCalc(e *echo.Echo) {
	e.PUT("/:id", h.updateCalc)
}

func (h *Handler) updateCalc(c echo.Context) (err error) {
	dto := new(CreateInfoDto)
	elem := new(Webhook)
	id := c.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	elem, err = h.StoreWebhook.Get(idInt64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !Equal(elem.Fields, dto.Fields) {
		elem.Fields = dto.Fields
	} else if elem.Op != dto.Op {
		elem.Op = dto.Op
	}

	if err := h.StoreWebhook.PutKey(elem, idInt64);
	err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, elem)
}