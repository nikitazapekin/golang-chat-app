


/*package router

import (
	"fmt"
	"github.com/nikita/go-microservices/controller"
	"github.com/labstack/echo/v4"
)

func Test(e *echo.Echo) {
	fmt.Println("hew")
}
func SetAuthRoutes(e *echo.Echo) {
 
	e.POST("/chat/sign-in", controller.Register)
	e.GET("/chat/test", Test(e *echo.Echo))
}
 */



 package router

import (
    "github.com/nikita/go-microservices/controller"
    "github.com/labstack/echo/v4"
	"net/http"
)

func Test(c echo.Context) error {
    return c.String(http.StatusOK, "Test")
}

func SetAuthRoutes(e *echo.Echo) {
    e.POST("/chat/sign-in", controller.Register)
    e.GET("/chat/test", Test)
	e.GET("/chat/token", controller.GetAccessToken)
	e.GET("/chat/refresh-token", controller.GetRefreshToken)
}
