package app

import (
	"project-layout/configs"
	"project-layout/internal/routes"

	"github.com/labstack/echo/v4"
)

func Run(){
	e:=echo.New()
	configs.Load()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(configs.Server.Address))
}
