package router

import (
	"github.com/nikita/go-microservices/controller"
	"github.com/labstack/echo/v4"
)

func SetOffersRoutes(e *echo.Echo) {
 
	e.POST("/worklist.com/getOffers", controller.Register)
	 
}