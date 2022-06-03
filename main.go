package main

import (
	"catchy/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for Users
	e.GET("/api/v1/user", handler.User)
	e.POST("/api/v1/user", handler.CreateUser)
	e.DELETE("/api/v1/user", handler.DeleteUser)
	e.PUT("/api/v1/user/:id", handler.UpdateUser)

	// Routes for chatrooms
	e.GET("/api/v1/chatroom", handler.ChatRoom)
	e.POST("/api/v1/chatroom", handler.CreateChatRoom)
	e.DELETE("/api/v1/chatroom", handler.DeleteChatRoom)
	e.PUT("/api/v1/chatroom/:id", handler.UpdateChatRoom)

	e.GET("/socket/:id", handler.WebsocHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
