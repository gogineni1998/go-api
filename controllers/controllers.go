package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gogineni1998/go-api/models"
	"github.com/gogineni1998/go-api/services"
	"github.com/gogineni1998/go-api/utilities"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := utilities.EstablishConnection()
	defer db.Close()
	users := []models.User{}
	services.GetUsers(&users, db)
	return c.JSON(http.StatusAccepted, users)
}

func GetUser(c echo.Context) error {
	user := models.User{}
	db := utilities.EstablishConnection()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	err = services.GetUser(id, &user, db)
	if err != nil {
		log.Println(err)
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
	db := utilities.EstablishConnection()
	defer db.Close()
	userId, err := services.CreateUser(user, db)
	if err != nil {
		response.Message = "Unable to create the user"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Message = "User " + strconv.Itoa(userId) + " Created Successfully"
	}
	return c.JSON(http.StatusOK, response)
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
	db := utilities.EstablishConnection()
	defer db.Close()
	_, err = services.UpdateUser(user, db)
	if err != nil {
		log.Println(err)
	}
	response.Message = "updated successfully"
	return c.JSON(http.StatusAccepted, response)
}

func DeleteUser(c echo.Context) error {
	response := new(models.Response)
	db := utilities.EstablishConnection()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	_, err = services.DeleteUser(id, db)
	if err != nil {
		log.Println(err)
	}
	response.Message = "deleated successfully"
	return c.JSON(http.StatusAccepted, response)
}
