package app

import (
	"go-weathermonitor/configs"
	"go-weathermonitor/internal/routes"

	"github.com/labstack/echo/v4"
)

func Run(){
	e:=echo.New()
	configs.Load()
	configs.Databaseinit()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(configs.Server.Address))
}
