package router

import (
	"github.com/labstack/echo/v4"
)
func InitRoutes(e *echo.Echo) {

	SetAuthRoutes(e)
	SetPersonalRoutes(e)
	 SetSearchRoutes(e)
}