package models

import (
	"github.com/gorilla/websocket"
)

//model for User
type Client struct {
	UserID     string
	Connection *websocket.Conn
}

//model for subscription
type Subscription struct {
	ChatroomId string
	Clients    *[]Client
}

//model for Message format
type Message struct {
	Action     string `json:"action"`
	ChatroomId string `json:"chatroomId"`
	Message    string `json:"message"`
}
