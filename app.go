package main

import (
	"karsingh991/cns-auth/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var app *echo.Echo

func start() {
	app = echo.New()
}

//register all the middlwares here
func registerMiddleware() {
	//logger middleware
	app.Use(middleware.Logger())
}

//register all the rest apis here
func registerRestApis() {
	app.GET("/", handlers.HealthHandler)
	http.HandleFunc("/user/create", createUserHandler)
}
