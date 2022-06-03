package handler

import (
	"catchy/models"
	"catchy/service"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var server = &service.Server{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func to serve websocket request
func WebsocHandler(c echo.Context) error {

	//Extracting UserId coming from request header
	id := c.Param("id")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// upgrades connection to websocket
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// create new User & add to Users list
	client := models.Client{
		UserID:     id,
		Connection: conn,
	}

	// welcome message for new users
	server.Send(&client, "Welcome to Catchy! Your Id is - "+client.UserID)

	// message handling unil server exit or disconnect from client
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			server.RemoveUser(client)
			return err
		}
		server.ProcessMessage(client, p)
	}
}
