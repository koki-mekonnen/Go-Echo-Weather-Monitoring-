package routes

import (
	"project-layout/internal/app/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo){
	e.GET("/",handlers.HomeHandler)
}
