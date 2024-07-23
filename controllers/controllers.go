package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gogineni1998/go-api/models"
	"github.com/gogineni1998/go-api/services"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []models.User{}
	services.GetUsers(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := models.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusBadRequest, user)
	}
	err = services.GetUser(id, &user)
	if err != nil {
		log.Default().Println(err)
		return c.JSON(http.StatusBadRequest, user)
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	response := new(models.Response)
	err := c.Bind(user)
	if err != nil {
		log.Default().Println(err)
		response.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, response)
	}
	created_user, err := services.CreateUser(user)
	if err != nil {
		log.Default().Println(err)
		response.Message = "Unable to create the user"
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, created_user)
}

func UpdateUser(c echo.Context) error {
	user := new(models.User)
	response := new(models.Response)
	err := c.Bind(user)
	if err != nil {
		log.Default().Println(err)
		response.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, response)
	}
	_, err = services.UpdateUser(user)
	if err != nil {
		log.Default().Println(err)
		response.Message = "Unable to update the user"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "updated successfully"
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	response := new(models.Response)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Default().Println(err)
		response.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, response)
	}
	_, err = services.DeleteUser(id)
	if err != nil {
		log.Default().Println(err)
		response.Message = "unable to deleated the user"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "deleated successfully"
	return c.JSON(http.StatusOK, response)
}
