package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"DGM-B/db"
	route "DGM-B/router"
)

func main() {
	db.Connect()
	e := echo.New()
	e.Use(middleware.CORS())
	route.RegisteRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
