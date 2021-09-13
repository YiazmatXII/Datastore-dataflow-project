package handler

import (
	. "app/computation"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetResult(t *testing.T) {
	c, _, rec := initEchoTest("/", nil, http.MethodPost)
	h, _, msc := initMockedHandler()
	c.SetPath("/:webhook_id/computation/:computation_id")
	c.SetParamNames("webhook_id", "computation_id")
	c.SetParamValues("5632499082330112", "2345673456745832")
	keys := []KeyValue{
		{
			Key: "a",
			Value: 5,
		},
		{
			Key: "b",
			Value: 7,
		},
	}
	result := Computation{
		ID: 		2345673456745832,
		WebhookID: 	5632499082330112,
		Values: keys,
		Result: 	12,
		Processed: true,
	}
	msc.On("Get", int64(2345673456745832)).Return(&result, nil)
	if assert.NoError(t, h.getResult(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":2345673456745832,"webhook_id":5632499082330112,"values":[{"key":"a","value":5},{"key":"b","value":7}],"result":12,"processed":true}` + "\n", rec.Body.String())
	}
}

func TestGetResultFails(t *testing.T) {
	c, _, _ := initEchoTest("/", nil, http.MethodPost)
	h, _, msc := initMockedHandler()
	keys := []KeyValue{
		{
			Key: "a",
			Value: 5,
		},
		{
			Key: "b",
			Value: 7,
		},
	}
	result := Computation{
		ID: 		2345673456745832,
		WebhookID: 	5632499082330112,
		Values: keys,
		Result: 	12,
		Processed: false,
	}

	msc.On("Get", int64(2345673456745832)).Return(&result, nil)
	if err := h.getResult(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}

	c.SetPath("/:webhook_id/computation/:computation_id")
	c.SetParamNames("webhook_id", "computation_id")
	c.SetParamValues("5632499082330112", "2345673456745832")
	if err := h.getResult(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, 404, herr.Code)
	}

	msc.On("Get", int64(2345673456745832)).Return(&result, errors.New("error"))
	if err := h.getResult(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
}