package models

//model for User
type User struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Username string `bson:"username"`
}
