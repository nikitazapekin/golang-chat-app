
package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"
)

func SearchUsers(c echo.Context) error {
    username := c.QueryParam("username")
    fmt.Println("Username:", username)

    users, err := db.FindUsersDataByUsername(username)
   // fmt.Println("USERS")
   // fmt.Println(users)
    
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{"users": users})
}
