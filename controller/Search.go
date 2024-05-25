
package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"

	//"encoding/json"
	//"github.com/dgrijalva/jwt-go"
	//"strings"
)




/*

func SearchUsers(c echo.Context) error {
	username := c.QueryParam("username")
	fmt.Println("Username:", username)


	fmt.Println("Username:", username)

	foundUsername, country, tel, tokenString, chats, avatar, describtion, err := db.FindUserDataByUsername(username)
	fmt.Println(foundUsername, country, tel, tokenString, chats, avatar, describtion, err)

	return c.JSON(http.StatusOK, map[string]string{"message": "Token parsed successfully"})
 
}
 */

 func SearchUsers(c echo.Context) error {
    username := c.QueryParam("username")
    fmt.Println("Username:", username)
    users, err := db.FindUsersDataByUsername(username)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, users)
}

