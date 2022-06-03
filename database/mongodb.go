package database

import (
	"catchy/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

//MongoDB database connection
func ConnectDb() {
	var err error
	//mongodb client connection
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	//connection check
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}

// function to get Allchatroom Data from DataBase
func GetChatRooms() ([]models.Chatroom, error) {
	//first time connection or if the connection is lost
	if client == nil {
		ConnectDb()
	}
	chatRoomCollection := client.Database("catchy").Collection("chatroom")
	cursor, err := chatRoomCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.Chatroom{}, err
	}
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []models.Chatroom{}, err
	}

	var chatrooms []models.Chatroom
	for _, result := range results {
		fmt.Println(result)
		data, err := bson.Marshal(result)
		if err != nil {
			return []models.Chatroom{}, err
		}
		var chatroom models.Chatroom
		err = bson.Unmarshal(data, &chatroom)
		if err != nil {
			return []models.Chatroom{}, err
		}
		chatrooms = append(chatrooms, chatroom)

	}
	return chatrooms, nil
}

// function to get Create New Chatroom in database
func CreateChatRoom(chatroom models.Chatroom) (models.Chatroom, error) {
	if client == nil {
		ConnectDb()
	}
	chatRoomCollection := client.Database("catchy").Collection("chatroom")

	id := primitive.NewObjectID()
	newChatroom := models.Chatroom{
		Id:          id.Hex(),
		Name:        chatroom.Name,
		Description: chatroom.Description,
	}
	result, err := chatRoomCollection.InsertOne(context.TODO(), newChatroom)
	if err != nil {
		return models.Chatroom{}, err
	}

	return GetChatRoomById(result.InsertedID), nil
}

// function to get specified Chatroom data from DB by passing ChatroomId
func GetChatRoomById(Id interface{}) models.Chatroom {
	if client == nil {
		ConnectDb()
	}

	//id := fmt.Sprint(Id)

	var chatroom models.Chatroom
	chatRoomCollection := client.Database("catchy").Collection("chatroom")

	err := chatRoomCollection.FindOne(context.TODO(), bson.M{"_id": Id}).Decode(&chatroom)
	if err != nil {
		return models.Chatroom{}
	}

	fmt.Println(chatroom.Name)
	id1 := primitive.NewObjectID()
	newChatroom := models.Chatroom{
		Id:          id1.Hex(),
		Name:        chatroom.Name,
		Description: chatroom.Description,
	}

	return newChatroom
}

// function to delete specified Chatroom from database by passing chatroomId
func DeleteChatRoom(id string) error {
	if client == nil {
		ConnectDb()
	}
	ctx := context.Background()
	chatRoomCollection := client.Database("catchy").Collection("chatroom")
	_, err := chatRoomCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
		return err
	}
	return err
}

// function to Update specified Chatroom data from database by passing ChatroomId and data to be modified
func UpdateChatRoom(id string, chatroom models.Chatroom) (models.Chatroom, error) {
	if client == nil {
		ConnectDb()
	}
	chatRoomCollection := client.Database("catchy").Collection("chatroom")
	// Id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"_id", bson.D{{"$eq", id}}},
				},
			},
		},
	}
	// create the update query
	update := bson.D{
		{"$set",
			bson.D{
				{"name", chatroom.Name},
			},
		},
	}

	result, err := chatRoomCollection.UpdateOne(context.TODO(), filter, update)
	// check for errors in the updating
	fmt.Println(result)
	if err != nil {
		return models.Chatroom{}, err
	}
	return GetChatRoomById(id), nil
}
