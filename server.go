//module github.com/nikita/go-microservices
/*

package main

import (
	//"fmt"
	"log"
	"net/http"
	 
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/nikita/go-microservices/db"
	r "github.com/nikita/go-microservices/routes"


	"github.com/gorilla/websocket" 
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	e := echo.New()
	db.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:3000"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002",  "http://localhost:5500"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	
	
	e.GET("/chat/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r.InitRoutes(e)
	
	//r.InitWebsocketRoutes(e)
	 
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

*/
/*
package main

import (
	"fmt"
	"log"
	"net/http"
	 
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/nikita/go-microservices/db"
	r "github.com/nikita/go-microservices/routes"


	"github.com/gorilla/websocket" 
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

 
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}



func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }
 
    reader(ws)
}
 

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}




func main() {
	e := echo.New()
	db.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:3000"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002",  "http://localhost:5500"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	
	
	e.GET("/chat/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r.InitRoutes(e)
	
	//r.InitWebsocketRoutes(e)
	setupRoutes()
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

*/






/*
"fmt"
	"log"
	"net/http"
	 
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/nikita/go-microservices/db"
	r "github.com/nikita/go-microservices/routes"

	*/




























	/*
package main

import (
    "fmt"
    "log"
    "net/http"
	"github.com/gorilla/websocket" 


	"github.com/labstack/echo/v4"
	db "github.com/nikita/go-microservices/db"
	r "github.com/nikita/go-microservices/routes"
	"github.com/labstack/echo/v4/middleware"
)


var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

 
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}



func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }
 
    reader(ws)
}
 

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	
	







	
	e := echo.New()
	db.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:3000"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002",  "http://localhost:5500"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	
	
	e.GET("/chat/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	r.InitRoutes(e)
	fmt.Println("Hello World")
	setupRoutes()

	
    log.Fatal(http.ListenAndServe(":5000", nil))
}
	
*/









package main

import (
	"fmt"
	"log"
	"net/http"

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

func reader(conn *websocket.Conn) {
	for {
		 
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		 
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }


	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

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
