package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	. "app/webhook"
)

func TestGetCalc(t *testing.T) {
	c, _, rec := initEchoTest("/", nil, http.MethodPost)
	h, msw, _ := initMockedHandler()
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("5632499082330112")
	var testFields = []string{"a", "b", "c"}
	elem := Webhook{
		ID: 5632499082330112,
		Fields: testFields,
		Op:     "sub",
	}
	msw.On("Get", int64(5632499082330112)).Return(&elem, nil)
	if assert.NoError(t, h.getCalc(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":5632499082330112,"fields":["a","b","c"],"op":"sub"}` + "\n", rec.Body.String())
	}
}

func TestGetCalcFails(t *testing.T) {
	c, _, _ := initEchoTest("/", nil, http.MethodDelete)
	h, msw, _ := initMockedHandler()
	c.SetPath("/:id")
	var testFields = []string{"a", "b", "c"}
	elem := Webhook{
		ID: 5632499082330112,
		Fields: testFields,
		Op:     "sub",
	}
	msw.On("Get", int64(5632499082330112)).Return(&elem, errors.New("error"))
	if err := h.deleteCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("5632499082330112")
	if err := h.getCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
}