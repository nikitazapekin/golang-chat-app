
package router

import (
    "github.com/nikita/go-microservices/controller"
    "github.com/labstack/echo/v4"
//	"net/http"
)
 

func SetPersonalRoutes(e *echo.Echo) {
    e.GET("/chat/personal", controller.PersonalData)
    e.GET("/chat/personal-username", controller.PersonalDataByUsername)
}


//personal-username