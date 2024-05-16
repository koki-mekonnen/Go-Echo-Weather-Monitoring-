package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)


func HomeHandler(c echo.Context) error{
	fmt.Println("hello")
	return c.String(http.StatusOK, "Server is Running")

}
