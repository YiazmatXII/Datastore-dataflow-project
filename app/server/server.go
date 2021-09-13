package server

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	. "app/config"
	. "app/handler"
)

var (
	hookStore 		StoreWebhook
	computeStore 	StoreComputation
)

func applyMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func InitHandlers(e *echo.Echo, hookStore StoreWebhook, computeStore StoreComputation) {
	h := Handler{StoreWebhook: hookStore, ComputeStore: computeStore}
	h.AddCreateCalc(e)
	h.AddGetCalc(e)
	h.AddGetAllCalc(e)
	h.AddUpdateCalc(e)
	h.AddDeleteCalc(e)
	h.AddGetResult(e)
	h.AddDeleteCalc(e)
}

func init() {
	ws, _ := NewDatastoreWebhookStore(DatastoreClient)
	hookStore = ws
	cs, _ := NewDatastoreComputationStore(DatastoreClient)
	computeStore = cs
}

func SetupEcho(e *echo.Echo) {
	applyMiddleware(e)
	InitHandlers(e, hookStore, computeStore)
	e.Validator = &CustomValidator{Validator: validator.New()}
}