  
-Chatroom Project Description

 Follow these setup to run the Project:

        1. Get the code from Github: 
                    $cd go/src
                    $ git clone https://github.com/Nagarjunhabbu/Catchy.git

    
        2. Install docker :
                - $ sudo apt install docker.io
                - $ sudo snap install docker
                - $ sudo docker images
                - $ sudo docker ps -a
                - $ sudo docker ps  (To check for containers in a running state, use the following command)

        3. Install mongoDB
                - $ sudo docker search mongodb
                - $ mkdir -pv mongodb/database
                - $ docker-compose up 


        4. Run the Server: $ make run
        5. Use any Websocket client for client server interaction
             - https://websocketking.com/  (websocket client)


----------------------------------------------------------------------------------------------------------------------------

All the CRUDL operations for users and chatrooms are performed using mongoDB
     
     To perform the CRUD operation use below link
        
  GET request-    http://localhost:8000/api/v1/user   
                 /user - to get all user data
                 /chatroom - to get all chatroom data

    
  POST Request -  http://localhost:8000/api/v1/user  - create new user
                    send JSON body along with this
                    {
                        "name":"abc"
                        "username":"abc12"
                    } 


  PUT Request -   http://localhost:8000/api/v1/user/{id}   - update specific user data
                    send Id in header and JSON body
                    {
                        "name":"xyz",
                    }
  
  DELETE Request -  http://localhost:8000/api/v1/user  - delete specific user data
                       send UserId in Json body
                       {
                        "id":"123",
                       }


---------------------------------------------------------------------------------------------------------------------------------

                                        Operations to be performed in websocketclient
                                                
step 1 : create two connection
step 2 : connect both users 

         ws://localhost:8000/socket/{user1_id}   
         ws://localhost:8000/socket/{user2_id}                                         

         click on connect on both sides                                                           
step 3 : subscribe to the chatroom from both side by giving same chatroomId
            {                                                                             
                "action" :"subscribe",                                                                                                         
                 "chatroomId" :"546",                                                            
                 "message":"Hello world!"                                                        
             }                                                                                 
    click send                                                                         
    both the users are connected to chatroom "546" so they can communicate until one of the user "unsubscribe"

step 4 : communicate by sending message from both Users 
                      {
                           "action" :"broadcast"
                            "chatroomId" :"546",
                            "message":"Hello world!"
                     }

step 4: Unsubscribe one user and send message from another user
            {
                    "action" :"unsubscribe",
                    "chatroomId" :"546",
                    "message":"Hello world!"
            }