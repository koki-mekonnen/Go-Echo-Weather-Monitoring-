package routes

import (
	"github.com/labstack/echo/v4"
	"go-weathermonitor/internal/app/handlers"

)

func RegisterRoutes(e *echo.Echo){
	e.GET("/:city",handlers.GetWeatherData)
}
