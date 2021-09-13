package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	. "app/webhook"
)

func TestGetAllCalc(t *testing.T) {
	c, _, rec := initEchoTest("/", nil, http.MethodGet)
	h, msw, _ := initMockedHandler()
	var testFields = []string{"a", "b", "c"}
	var list  = []*Webhook {
		{
			ID: 5632499082330112,
			Fields: testFields,
			Op:     "sub",
		},
		{
			ID: 5632499082330112,
			Fields: testFields,
			Op:     "sub",
		},
	}
	msw.On("GetAll").Return(list, nil)
	if assert.NoError(t, h.getAllCalc(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `[{"id":5632499082330112,"fields":["a","b","c"],"op":"sub"},{"id":5632499082330112,"fields":["a","b","c"],"op":"sub"}]` + "\n", rec.Body.String())
	}
}

func TestGetAllCalcFails(t *testing.T) {
	c, _, _ := initEchoTest("/", nil, http.MethodGet)
	h, msw, _ := initMockedHandler()
	var list []*Webhook
	msw.On("GetAll").Return(list, errors.New("error"))
	if err := h.getAllCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
}