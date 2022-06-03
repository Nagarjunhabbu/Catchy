package models

//model for chatroom
type Chatroom struct {
	Id          string `bson:"_id"`
	Name        string
	Description string
}
