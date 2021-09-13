package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestDeleteCalc(t *testing.T) {
	c, _, _ := initEchoTest("/", nil, http.MethodDelete)
	h, msw, _ := initMockedHandler()
	msw.On("Delete", mock.Anything).Return(nil)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("5632499082330112")
	assert.Nil(t, h.deleteCalc(c))
}

func TestDeleteCalcFails(t *testing.T) {
	c, _, _ := initEchoTest("/", nil, http.MethodDelete)
	h, msw, _ := initMockedHandler()
	msw.On("Delete", mock.Anything).Return(errors.New("error"))
	if err := h.deleteCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("5632499082330112")
	if err := h.deleteCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
}