package service

import (
	"catchy/models"
	"encoding/json"
)

//model for  server
type Server struct {
	Subscriptions []models.Subscription
}

//Sends message to single client
func (s *Server) Send(client *models.Client, message string) {
	client.Connection.WriteMessage(1, []byte(message))
}

//function to remove user from chatroom
func (s *Server) RemoveUser(client models.Client) {
	// Read all subs
	for _, sub := range s.Subscriptions {
		// Read all client
		for i := 0; i < len(*sub.Clients); i++ {
			if client.UserID == (*sub.Clients)[i].UserID {
				// If found, remove client
				if i == len(*sub.Clients)-1 {
					// if it's stored as the last element, crop the array length
					*sub.Clients = (*sub.Clients)[:len(*sub.Clients)-1]
				} else {
					// if it's stored in between elements, overwrite the element and reduce iterator to prevent out-of-bound
					*sub.Clients = append((*sub.Clients)[:i], (*sub.Clients)[i+1:]...)
					i--
				}
			}
		}
	}
}

// function to process user action
func (s *Server) ProcessMessage(client models.Client, payload []byte) *Server {
	m := models.Message{}
	if err := json.Unmarshal(payload, &m); err != nil {
		s.Send(&client, "Server: Invalid payload")
	}

	switch m.Action {
	case "broadcast":
		s.Broadcast(m.ChatroomId, []byte(m.Message), client)

	case "subscribe":
		s.Subscribe(&client, m.ChatroomId)
		s.Send(&client, "Subscribed Successfully!")

	case "unsubscribe":
		s.Unsubscribe(&client, m.ChatroomId)
		s.Send(&client, "Unsubscribed Successfully!")

	default:
		s.Send(&client, "Server: Action unrecognized")

	}

	return s
}

// function to check wether the user already exists in the chatrrom or not
func IsUserExists(client models.Client, clients []models.Client) bool {
	for _, val := range clients {
		if client == val {
			return true
		}
	}
	return false
}

// function to send messages to chatroom
func (s *Server) Broadcast(ChatroomId string, message []byte, client models.Client) {
	var clients []models.Client

	// get list of clients subscribed to topic
	for _, sub := range s.Subscriptions {
		if sub.ChatroomId == ChatroomId {
			clients = append(clients, *sub.Clients...)
		}

	}
	if !IsUserExists(client, clients) {
		client.Connection.WriteMessage(1, []byte("You cannot send message to this chatroom!!"))
		return
	}

	// sends the message to all the clients
	for _, client := range clients {
		s.Send(&client, string(message))
	}
}

// function to create chatroom if its new chatroom and adds user who creates it or adding Newuser if chatroom already exists
func (s *Server) Subscribe(client *models.Client, chatroomId string) {
	exist := false

	// check for chatroom existance
	for _, sub := range s.Subscriptions {
		// if found, add User
		if sub.ChatroomId == chatroomId {
			exist = true
			if IsUserExists(*client, *sub.Clients) {
				return
			}
			*sub.Clients = append(*sub.Clients, *client)
		}
	}

	// else, add new chatroom & add user to that chatroom
	if !exist {
		newUser := &[]models.Client{*client}

		newChatRoom := &models.Subscription{
			ChatroomId: chatroomId,
			Clients:    newUser,
		}

		s.Subscriptions = append(s.Subscriptions, *newChatRoom)
	}
}

func (s *Server) Unsubscribe(client *models.Client, chatroomId string) {
	// Read all cahtrooms
	for _, sub := range s.Subscriptions {
		if sub.ChatroomId == chatroomId {

			// Read all Chatroom's User
			for i := 0; i < len(*sub.Clients); i++ {
				if client.UserID == (*sub.Clients)[i].UserID {
					// If found, remove user
					if i == len(*sub.Clients)-1 {
						// if it's stored as the last element, just slice the last entry
						*sub.Clients = (*sub.Clients)[:len(*sub.Clients)-1]
					} else {
						// if it's stored in between elements, overwrite the element and reduce iterator to prevent out-of-bound
						*sub.Clients = append((*sub.Clients)[:i], (*sub.Clients)[i+1:]...)
						i--
					}
				}
			}
		}
	}
}
