
/*
package router

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
//	"log"
	"net/http"

)
  
 

func handleWebSocket(c echo.Context) error {
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true // Разрешаем все origin
        },
    }

    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    fmt.Println("Пользователь подключился")

    for {
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("Ошибка чтения сообщения:", err)
            break
        }

        fmt.Println("Получено сообщение от пользователя:", string(message))

        // Эхо-отправка сообщения обратно клиенту
        err = ws.WriteMessage(messageType, message)
        if err != nil {
            fmt.Println("Ошибка отправки сообщения:", err)
            break
        }
    }

    fmt.Println("Пользователь отключился")
    return nil
}

func InitWebsocketRoutes(e *echo.Echo) {
    e.GET("/", handleWebSocket)
}
*/










/*

package router

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

func handleWebSocket(c echo.Context) error {
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true // Allow all origins
        },
    }

    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    // Display a message when a user connects
    fmt.Println("User connected")

    for {
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }

        fmt.Println("Received message from user:", string(message))

        // Echo the message back to the client
        err = ws.WriteMessage(messageType, message)
        if err != nil {
            fmt.Println("Error sending message:", err)
            break
        }
    }

    // Display a message when a user disconnects
    fmt.Println("User disconnected")
    return nil
}

func InitWebsocketRoutes(e *echo.Echo) {
    e.GET("/", handleWebSocket)
}

*/





















/*

package router

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        // Разрешаем только запросы с определенного origin (http://localhost:5500)
        return r.Header.Get("Origin") == "http://localhost:5500"
    },
}

func handleWebSocket(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    // Отправляем сообщение о подключении
    err = ws.WriteMessage(websocket.TextMessage, []byte("User connected"))
    if err != nil {
        fmt.Println("Error sending message:", err)
        return err
    }

    for {
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }

        fmt.Println("Received message from user:", string(message))

        // Отправляем сообщение обратно клиенту
        err = ws.WriteMessage(messageType, message)
        if err != nil {
            fmt.Println("Error sending message:", err)
            break
        }
    }

    // Отображаем сообщение при отключении пользователя
    fmt.Println("User disconnected")
    return nil
}

func InitWebsocketRoutes(e *echo.Echo) {
    e.GET("/", handleWebSocket)
}
*/




















package router

import (
    "fmt"
    "log"
    "net/http"
	"github.com/gorilla/websocket" 
	"github.com/labstack/echo/v4"
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
func InitWebsocketRoutes(e *echo.Echo) {
	http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}


/*
func handleWebSocket() {
    fmt.Println("Hello World")
    setupRoutes()
  //  log.Fatal(http.ListenAndServe(":8080", nil))
}
	
*/










/*


package router

import (
	"fmt"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

func handleWebSocket(c echo.Context) error {
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true // Allow all origins
        },
    }

    // Подключение к WebSocket
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    // Оповещение о подключении пользователя
    fmt.Println("Пользователь подключился")

    for {
        // Чтение сообщения от клиента
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("Ошибка чтения сообщения:", err)
            break
        }

        // Вывод принятого сообщения от клиента
        fmt.Println("Получено сообщение от пользователя:", string(message))

        // Эхо-отправка сообщения обратно клиенту
        err = ws.WriteMessage(messageType, message)
        if err != nil {
            fmt.Println("Ошибка отправки сообщения:", err)
            break
        }
    }

    // Оповещение о отключении пользователя
    fmt.Println("Пользователь отключился")
    return nil
}

// Инициализация маршрутов WebSocket
func InitWebsocketRoutes(e *echo.Echo) {
    e.GET("/ws", handleWebSocket) // WebSocket-маршрут на "/ws"
}

*/