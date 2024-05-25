
package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"

	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func ParseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return token, nil
}
 



func PersonalData(c echo.Context) error {
	username := c.QueryParam("username")
	fmt.Println("Username:", username)

	// Retrieve token string from Authorization header
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization header is missing or invalid"})
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the JSON to extract the token string
	var tokenJSON struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal([]byte(tokenString), &tokenJSON); err != nil {
		fmt.Println("Error decoding token JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token format"})
	}


	tokenString = tokenJSON.Token


	secretKey := []byte("your-secret-key")
	token, err := ParseToken(tokenString, secretKey)
	if err != nil {

		fmt.Println("Error parsing token:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse token"})
	}

	if token == nil {
		fmt.Println("Token is nil")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}


	fmt.Println("TOKEN:", token)

	usernameClaim, ok := token.Claims.(jwt.MapClaims)["username"]
	if !ok {
		fmt.Println("Username claim not found in token")
	
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Username claim not found in token"})
	}


	username, ok = usernameClaim.(string)
	if !ok {
		fmt.Println("Invalid username claim in token")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid username claim in token"})
	}

	fmt.Println("Username:", username)

	foundUsername, country, tel, tokenString, chats, avatar, describtion, err := db.FindUserDataByUsername(username)
	fmt.Println(foundUsername, country, tel, tokenString, chats, avatar, describtion, err)

	//return c.JSON(http.StatusOK, map[string]string{"message": "Token parsed successfully"})
	responseData := map[string]interface{}{
        "access_token": "newAccessToken",
        "username":     foundUsername,
        "country":      country,
        "tel":          tel,
        "chats":        chats,
        "avatar":       avatar,
        "description": " description",
    }

    return c.JSON(http.StatusOK, responseData)
}
 





















func PersonalDataByUsername(c echo.Context) error {
	username := c.QueryParam("user")
	fmt.Println("Username:", username)
	foundUsername, country, tel, refreshToken, chats, avatar, description, err := db.FindUserDataByUsername(username)
	fmt.Println(foundUsername, country, tel, refreshToken, chats, avatar, description, err)

if err!= nil {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Username claim not found in token"})
}
	responseData := map[string]interface{}{
        "access_token": "newAccessToken",
        "username":     foundUsername,
        "country":      country,
        "tel":          tel,
        "chats":        chats,
        "avatar":       avatar,
        "description": " description",
    }

    return c.JSON(http.StatusOK, responseData)
}
 




