package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) AddDeleteCalc(e *echo.Echo) {
	e.DELETE("/:id", h.deleteCalc)
}

func (h *Handler) deleteCalc(c echo.Context) (err error) {
	dto := new(GetInfoDto)
	if err = c.Bind(dto); err != nil {
		return err
	}
	if err = c.Validate(dto); err != nil {
		return err
	}

	if err := h.StoreWebhook.Delete(dto.ID);
		err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
