package api

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (api *API) Router() {
	e := echo.New()
	defineConfiguration(e)
	api.defineRoutes(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func defineConfiguration(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status} \n",
	}))
}

func (api *API) defineRoutes(e *echo.Echo) {
	e.GET("/invitation/all", api.GetInvitations)
	e.POST("/invitation", api.CreateInvitation)
}
