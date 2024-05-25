
package router

import (
    "github.com/nikita/go-microservices/controller"
    "github.com/labstack/echo/v4"
)
 

func SetSearchRoutes(e *echo.Echo) {
    e.GET("/chat/search-user", controller.SearchUsers)
 
}
