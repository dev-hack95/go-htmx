package main

import (
	"github.com/dev-hack95/go-htmx/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	userHandler := handlers.UserHanler{}
	router.GET("/user", userHandler.HandleUserShow)

	router.Start("0.0.0.0:8080")
}
