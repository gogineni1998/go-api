package controllers

import (
	"fmt"
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
	utilities.ErrorHanler(err)
	err = services.GetUser(id, &user, db)
	utilities.ErrorHanler(err)
	return c.JSON(http.StatusAccepted, user)
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
		fmt.Println(user)
	} else {
		response.Message = "User " + strconv.Itoa(userId) + " Created Successfully"
	}
	return c.JSON(http.StatusCreated, response)
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
	id, err := services.UpdateUser(user, db)
	utilities.ErrorHanler(err)
	if id == 0 {
		response.Message = "unable to update"
		return c.JSON(http.StatusConflict, response)
	}
	response.Message = "updated successfully"
	return c.JSON(http.StatusAccepted, response)
}

func DeleteUser(c echo.Context) error {
	response := new(models.Response)
	db := utilities.EstablishConnection()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	utilities.ErrorHanler(err)
	count, err := services.DeleteUser(id, db)
	utilities.ErrorHanler(err)
	if count == 0 {
		response.Message = "unable to deleate"
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Message = "deleated successfully"
	return c.JSON(http.StatusAccepted, response)
}
