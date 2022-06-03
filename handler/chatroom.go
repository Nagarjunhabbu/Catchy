package handler

import (
	"catchy/database"
	"catchy/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Handler function to GetAllChatroom data from databse
func ChatRoom(c echo.Context) error {
	chatrooms, err := database.GetChatRooms()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, chatrooms)
}

//Handler function for Chatroom creation
func CreateChatRoom(c echo.Context) error {
	var chatroom models.Chatroom
	if err := c.Bind(&chatroom); err != nil {
		return err
	}

	chatrooms, err := database.CreateChatRoom(chatroom)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, chatrooms)
}

//Handler function for Deleting Chatroom by passing ChatroomId
func DeleteChatRoom(c echo.Context) error {
	var chatroom models.Chatroom
	if err := c.Bind(&chatroom); err != nil {
		return err
	}
	err := database.DeleteChatRoom(chatroom.Id)
	if err == nil {
		return c.JSON(http.StatusOK, "Deleted Successfully!")
	}
	log.Error(err)
	return c.JSON(http.StatusInternalServerError, "Some Error Occurred while deleting doc! please try again")

}

//Handler function for Updating Chatroom data by passing ChatroomId and data to be modified

func UpdateChatRoom(c echo.Context) error {
	var chatroom models.Chatroom
	if err := c.Bind(&chatroom); err != nil {
		return err
	}
	id := c.Param("id")
	chatrooms, err := database.UpdateChatRoom(id, chatroom)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Oops!Try Again")
	}

	return c.JSON(http.StatusOK, chatrooms)
}
