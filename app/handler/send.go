package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	. "app/webhook"
)

func (h *Handler) AddSendCalc(e *echo.Echo) {
	e.GET("/:id/computation", h.sendCalc)
}

func (h *Handler) generatePayload(c echo.Context) (payload *Payload, err error) {
	var mapValues map[string]interface{}
	err = json.NewDecoder(c.Request().Body).Decode(&mapValues)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := c.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	elem := new(Webhook)

	elem, err = h.StoreWebhook.Get(idInt64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listValue := mapToKeyValue(mapValues)
	result, err := h.createComputation(idInt64, listValue)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	payload = &Payload{
		ComputationID: result.ID,
		WebhookID: 	   idInt64,
		Values:        listValue,
		Fields:        elem.Fields,
		Op:            elem.Op,
	}
	return payload, nil
}

func (h *Handler) sendCalc(c echo.Context) error {
	payload, err := h.generatePayload(c)
	if err != nil {
		return err
	}
	go publishOnTopic(*payload)
	return c.JSON(200, *payload)
}