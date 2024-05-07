package router

import (
	"github.com/labstack/echo/v4"
)
func InitRoutes(e *echo.Echo) {

	SetAuthRoutes(e)
	SetPersonalRoutes(e)
	/*SetAuth(e)
	SetBookRoutes(e)
	SetOffersRoutes(e)
	SetPersonalInformation(e)
	SetVacancy(e)
	SetGetPhotos(e)
	SetFilteredRoutes((e))
	SetGetUsers(e) */
}