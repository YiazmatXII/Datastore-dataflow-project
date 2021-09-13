package main

import (
	. "app/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	SetupEcho(e)
	e.Logger.Fatal(e.Start(":1323"))
}
