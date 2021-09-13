package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCreateCalc(t *testing.T) {
	var testString = strings.NewReader(`{"Fields": ["a", "b", "c"],"Op":"sub"}`)
	c, _, rec := initEchoTest("/", testString, http.MethodPost)
	h, msw, _ := initMockedHandler()
	msw.On("Put", mock.Anything).Return(nil)
	if assert.NoError(t, h.createCalc(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":345678945678,"fields":["a","b","c"],"op":"sub"}` + "\n", rec.Body.String())
	}
}

func TestCreateCalcFails(t *testing.T) {
	var testString = strings.NewReader(`{"Op":"add"}`)
	c, req, _ := initEchoTest("/", testString, http.MethodPost)
	h, msw, _ := initMockedHandler()
	msw.On("Put", mock.Anything).Return(errors.New("error"))
	if err := h.createCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
	req.Body = ioutil.NopCloser(strings.NewReader(`{"Fields": ["a", "b", "c"]}`))
	if err := h.createCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
	req.Body = ioutil.NopCloser(strings.NewReader(`"Fields": ["a", "b", "c"],"Op":"sub"}`))
	if err := h.createCalc(c); assert.Error(t, err) {
		herr, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusBadRequest, herr.Code)
	}
}