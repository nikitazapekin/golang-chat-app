

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	db "github.com/nikita/go-microservices/db"
	m "github.com/nikita/go-microservices/routes"

	_ "github.com/lib/pq"  
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
 
type Message struct {
	Name    string
	Message string
	To      string
	From    string
}
var (
	connections = make(map[string]*websocket.Conn)
	mu          sync.Mutex
)


func reader(conn *websocket.Conn, userName string) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			mu.Lock()
			delete(connections, userName)
			mu.Unlock()
			return
		}

		fmt.Println("MES TYPE:", messageType)
		fmt.Println("Raw message:", string(p))

		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		fmt.Printf("Parsed message: Name=%s, Message=%s, To=%s\n", msg.Name, msg.Message, msg.To)

		foundUsernameFrom, countryFrom, telFrom, refreshTokenFrom, chatsFrom, avatarFrom, descriptionFrom, errFrom := db.FindUserDataByUsername(msg.Name)
		fmt.Println("WEB SOCKETSSSSSSSSS",foundUsernameFrom, countryFrom, telFrom, refreshTokenFrom, chatsFrom, avatarFrom, descriptionFrom, errFrom)
		foundUsernameTo, countryTo, telTo, refreshTokenTo, chatsTo, avatarTo, descriptionTo, errTo := db.FindUserDataByUsername(msg.To)
		fmt.Println(foundUsernameTo, countryTo, telTo, refreshTokenTo, chatsTo, avatarTo, descriptionTo, errTo)

 
	if msg.Name != "" && msg.Message != "" && msg.To != "" {
		fmt.Println("ADDDDING")
		db.AddMessageToChatsTable(msg.Name, msg.Message, msg.To)
		db.AddMessageToGetterChatsTable(msg.Name, msg.Message, msg.To)
	} else {
		fmt.Println("Error: One or more message fields are empty.", "fr", msg.Name, "mes", msg.Message, "to", msg.To)
	}
		if msg.To != "" {
			mu.Lock()
			targetConn, ok := connections[msg.To]
			mu.Unlock()
			if ok {
				if err := targetConn.WriteMessage(messageType, p); err != nil {
					log.Println("Error writing message:", err)
				}
			} else {
				log.Printf("User %s not connected\n", msg.To)
			}
		} else {
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println("Error writing message:", err)
				return
			}
		}
	}
}

func wsEndpoint(c echo.Context) error {
	userName := c.QueryParam("user")
	if userName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User name is required")
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("USER IS CONNECTED:", userName)
	log.Println("Client Connected")

	mu.Lock()
	connections[userName] = ws
	mu.Unlock()

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
		return err
	}

	reader(ws, userName)
	return nil
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Home Page")
}











func readerChat(conn *websocket.Conn, userName string) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			mu.Lock()
			delete(connections, userName)
			mu.Unlock()
			return
		}

		fmt.Println("MES TYPE:", messageType)
		fmt.Println("Raw message:", string(p))

		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		fmt.Printf("Parsed message: Name=%s, Message=%s, To=%s\n", msg.Name, msg.Message, msg.To)
	}

}

func wsChat(c echo.Context) error {
	userName := c.QueryParam("user")
	companion := c.QueryParam("companion")
	fmt.Println("COMP", companion)
	if userName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User name is required")
	}
fmt.Println("WSSSSSSSSSSSSSSSSSSSSSS CHAAAAAAAAAAAAAAAAT")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("USER IS CONNECTED:", userName)
	log.Println("Client Connected")

	mu.Lock()
	connections[userName] = ws
	mu.Unlock()

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
		return err
	}
	//chats, err := db.FindUsersChat(userName)
	db.FindUsersChat(userName, companion)
	//fmt.Println(chats)
//	readerChat(ws, userName)
	return nil
}









func main() {
	r := echo.New()
	db.Connect()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "http://localhost:5500"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	r.GET("/chat/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r.GET("/", homePage)
	r.GET("/ws", wsEndpoint)

	r.GET("/ws/chat", wsChat)
	m.InitRoutes(r)
	fmt.Println("Hello World")

	log.Fatal(r.Start(":5000"))
}

/*


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>WebSocket Client</title>
</head>
<body>
    <h2>WebSocket Client</h2>
    <label for="userName">Username:</label>
    <input type="text" id="userName" placeholder="Enter your username">
    <button onclick="connect()">Connect</button>

    <div id="status"></div>

    <h3>Send Message</h3>
    <label for="toUser">To:</label>
    <input type="text" id="toUser" placeholder="Recipient username">
    <label for="message">Message:</label>
    <input type="text" id="message" placeholder="Enter your message">
    <button onclick="sendMessage()">Send</button>

    <h3>Messages</h3>
    <div id="messages"></div>

    <script>
        let ws;

        function connect() {
            const userName = document.getElementById('userName').value;
            if (!userName) {
                alert('Username is required');
                return;
            }

            const url = `ws://localhost:5000/ws?user=${userName}`;
            ws = new WebSocket(url);

            ws.onopen = () => {
                document.getElementById('status').innerText = 'Connected';
                console.log('Connected to the server');
            };

            ws.onmessage = (event) => {
                const messageDiv = document.createElement('div');
                messageDiv.textContent = `Received: ${event.data}`;
                document.getElementById('messages').appendChild(messageDiv);
                console.log('Received message:', event.data);
            };

            ws.onclose = () => {
                document.getElementById('status').innerText = 'Disconnected';
                console.log('Disconnected from the server');
            };

            ws.onerror = (error) => {
                console.log('WebSocket error:', error);
            };
        }

        function sendMessage() {
            const toUser = document.getElementById('toUser').value;
            const message = document.getElementById('message').value;
            const userName = document.getElementById('userName').value;

            if (!ws || ws.readyState !== WebSocket.OPEN) {
                alert('WebSocket connection is not open');
                return;
            }

            if (!toUser || !message) {
                alert('To and Message fields are required');
                return;
            }

            const msg = {
                name: userName,
                message: message,
                to: toUser
            };

            ws.send(JSON.stringify(msg));
            console.log('Sent message:', msg);
        }
    </script>
</body>
</html>



*/

// go run server.go
// go mod init github.com/nikita/go-microservices
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get -u github.com/golang/protobuf/protoc-gen
//go get -u google.golang.org/grpc
// mkdir chat
//go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
