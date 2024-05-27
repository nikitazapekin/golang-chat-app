
package main
import (
	"fmt"
	"log"
	"net/http"
		"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"


m"github.com/nikita/go-microservices/routes"
db "github.com/nikita/go-microservices/db"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
/*
func reader(conn *websocket.Conn) {
	for {
		 
		messageType, p, err := conn.ReadMessage()
		fmt.Println("MES TYPE")
		fmt.Println(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("PPPPPPPPPPPPPPPPPPPPPPPPPPPPPp")
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
*/




type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	To      string `json:"to"`
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
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

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}
/*
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		fmt.Println("MES TYPE:", messageType)
		fmt.Println("Raw message:", string(p))



		
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}
 */
func wsEndpoint(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }


	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
fmt.Println("USER IS CONNECTED")
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
		return err
	}

	reader(ws)
	return nil
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Home Page")
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
	m.InitRoutes(r)
	fmt.Println("Hello World")

	log.Fatal(r.Start(":5000"))
}



// go run server.go
// go mod init github.com/nikita/go-microservices
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get -u github.com/golang/protobuf/protoc-gen
//go get -u google.golang.org/grpc
// mkdir chat
//go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
