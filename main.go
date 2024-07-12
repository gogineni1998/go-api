package main

import (
	"github.com/gogineni1998/go-api/controllers"
	"github.com/gogineni1998/go-api/utilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	utilities.CreateTable()
	e.Start("0.0.0.0:8081")
}
