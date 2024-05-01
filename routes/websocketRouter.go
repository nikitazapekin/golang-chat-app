package router

import (
	//"fmt"
	//"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
//	"log"
//	"net/http"


	//"encoding/json"
 //   "time"
	//e "server/middleware"
//	m "server/db"
)
/*
var numberOfOnlineUsers = 0
var clients = make(map[*websocket.Conn]bool)
var onlineUsersMessage = "%d"
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleWebSocket(c echo.Context) error { //ХЭНДЛЕРРРРРРРРРРРРРРРР
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	clients[conn] = true
	numberOfOnlineUsers++
	broadcastOnlineUsersMessage()
	
	defer func() {
		numberOfOnlineUsers--
		delete(clients, conn)
		broadcastOnlineUsersMessage()
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Printf("Received message: %s\n", p)
		fmt.Println("CONNNN")
		fmt.Println(conn)
		fmt.Println(messageType)
		fmt.Println(p)
		err = conn.WriteMessage(messageType, []byte("Hello, client!"))
		if err != nil {
			return err
		}
	}
}
func broadcastOnlineUsersMessage() {
	message := fmt.Sprintf(onlineUsersMessage, numberOfOnlineUsers)
	fmt.Println("MESSSSSSSSSAAAGE")
	fmt.Println(message)
	fmt.Println(clients)
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
		}
	}
}

 


type Message struct {
    Username string `json:"username"`
    Token    string `json:"token"`
    Message  string `json:"message"`
}

func handlePublicChat(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    clients[conn] = true

    defer func() {
        delete(clients, conn)
        conn.Close()
    }()

    _, msgBytes, err := conn.ReadMessage()
    if err != nil {
        log.Printf("Error reading message from client: %v", err)
        return err
    }

    var incomingMessage Message
    err = json.Unmarshal(msgBytes, &incomingMessage)
    if err != nil {
        log.Printf("Error unmarshaling message: %v", err)
        return err
    }

 token := incomingMessage.Token
 decodedToken, errToken := e.Decode(token, "key")
 fmt.Println(errToken)  

 fmt.Println("Decoded token")
 fmt.Println(decodedToken)
 user,err := m.FindUserByUsername(decodedToken.Username)
 fmt.Println("USERrr")
 fmt.Println(user)
 fmt.Println("AVATARRRRRRRRRRRRRR")
fmt.Println(user.Avatar)
avatarFilename := user.Avatar
avatarURL := fmt.Sprintf("http://localhost:5000/worklist.com/image/%s", avatarFilename)
fmt.Println(avatarURL)
    currentTime := time.Now().Format(time.RFC3339)
    newMessage := struct {
        Username string `json:"username"`
        Message  string `json:"message"`
        Data     string `json:"data"`
		Avatar   string `json:"avatar"`
    }{
        Username: incomingMessage.Username,
        Message:  incomingMessage.Message,
        Data:     currentTime,
		Avatar:  avatarFilename,
		//Avatar:   avatarURL,
    }

    newMessageBytes, err := json.Marshal(newMessage)
    if err != nil {
        log.Printf("Error marshaling new message: %v", err)
        return err
    }

    handlePublicMessages(string(newMessageBytes))

    return nil
}

func handlePublicMessages(message string) error {
    fmt.Println("Received message:", message)
    fmt.Println(clients)
    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Printf("Error sending message to client: %v", err)
        }
    }
    return nil
}
 */
func InitWebsocketRoutes(e *echo.Echo) {
	//e.GET("/ws", handleWebSocket)
//	e.GET("/publicMessages", handlePublicChat)

}