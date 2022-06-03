package database

import (
	"catchy/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to get AllUserData from DataBase
func GetUser() ([]models.User, error) {

	//first time connection or if the connection is lost
	if client == nil {
		ConnectDb()
	}
	userCollection := client.Database("catchy").Collection("users")
	cursor, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.User{}, err
	}
	var results []bson.M

	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []models.User{}, err
	}

	var users []models.User
	for _, result := range results {
		fmt.Println(result)
		data, err := bson.Marshal(result)
		if err != nil {
			return []models.User{}, err
		}
		var user models.User
		err = bson.Unmarshal(data, &user)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)

	}
	return users, nil
}

// function to get Create New User in database
func CreateUser(user models.User) (models.User, error) {
	if client == nil {
		ConnectDb()
	}
	userCollection := client.Database("catchy").Collection("users")

	id := primitive.NewObjectID()
	newUser := models.User{
		Id:       id.Hex(),
		Name:     user.Name,
		Username: user.Username,
	}
	result, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return models.User{}, err
	}
	return GetUserById(result.InsertedID), nil
}

// function to get specified user data from DB by passing userId
func GetUserById(Id interface{}) models.User {
	if client == nil {
		ConnectDb()
	}

	var user models.User
	userCollection := client.Database("catchy").Collection("users")
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": Id}).Decode(&user)
	if err != nil {
		return models.User{}
	}

	userId := primitive.NewObjectID()
	newUser := models.User{
		Id:       userId.Hex(),
		Name:     user.Name,
		Username: user.Username,
	}
	return newUser
}

// function to delete specified user from database by passing userId
func DeleteUser(id string) error {
	if client == nil {
		ConnectDb()
	}
	ctx := context.Background()
	userCollection := client.Database("catchy").Collection("users")
	_, err := userCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
		return err
	}
	return nil
}

// function to Update specified user data from database by passing userId and data to be modified
func UpdateUser(id string, user models.User) (models.User, error) {
	if client == nil {
		ConnectDb()
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	userCollection := client.Database("catchy").Collection("users")
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"_id", bson.D{{"$eq", objID}}},
				},
			},
		},
	}
	// create the update query
	update := bson.D{
		{"$set",
			bson.D{
				{"name", user.Name},
			},
		},
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	// check for errors in the updating
	if err != nil {
		return models.User{}, err
	}
	return GetUserById(objID), nil
}
