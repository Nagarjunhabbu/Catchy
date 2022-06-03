package handler

import (
	"catchy/database"
	"catchy/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// function to GetAllUser data from databse
func User(c echo.Context) error {
	users, err := database.GetUser()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

//Handler function for User creation
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	users, err := database.CreateUser(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

//Handler function for Deleting User by passing userId
func DeleteUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	err := database.DeleteUser(user.Id)
	if err == nil {
		return c.JSON(http.StatusOK, "Deleted Successfully!")
	}
	return c.JSON(http.StatusOK, err)

}

///Handler function for Updating User data by passing userId and data to be modified

func UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	id := c.Param("id")
	users, err := database.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	return c.JSON(http.StatusOK, users)
}
