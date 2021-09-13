package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	. "app/computation"
)

func (h *Handler) AddGetResult(e *echo.Echo) {
	e.GET("/:webhook_id/computation/:computation_id", h.getResult)
}

func (h *Handler) getResult(c echo.Context) (err error) {
	dto := new(GetResultDto)
	if err = c.Bind(dto); err != nil {
		return err
	}
	if err = c.Validate(dto); err != nil {
		return err
	}

	result := new(Computation)

	result, err = h.ComputeStore.Get(dto.ComputationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.Processed == false {
		return echo.NewHTTPError(404, "Result not calculated")
	}

	return c.JSON(http.StatusOK, result)
}