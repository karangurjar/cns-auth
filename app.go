package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var app *echo.Echo

func Start() {
	app = echo.New()
}

func registerMiddleware() {
	app.Use(middleware.Logger())
}

func registerRestApis() {
	app.GET("/", healthHandler)
	app.Post("/user/create", createUserHandler)
	app.Get("/user", getUserHandler)
}
