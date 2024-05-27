

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
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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

