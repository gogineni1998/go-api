package routes

import (
	"net/http"

	"github.com/gogineni1998/go-api/controllers"
	"github.com/gogineni1998/go-api/utilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	utilities.CreateTable()
	return e
}
