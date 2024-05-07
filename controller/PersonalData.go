/*
package controller

import (
	//"encoding/json"
	"fmt"
	"net/http"
	//"time"

	"github.com/labstack/echo/v4"

	"github.com/nikita/go-microservices/db"
	//"github.com/dgrijalva/jwt-go"
)


func PersonalData(c echo.Context) error {
	fmt.Println("personal")
//(&foundUsername, &country, &tel, &token, &chats, &avatar, &describtion)
	username, country, tel, token, chats, avatar, describtion, err :=db.FindUserDataByUsername("nek")

	fmt.Println(username, country, tel, token, chats, avatar, describtion, err)

	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})


}
*/

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

/*
func PersonalData(c echo.Context) error {
	username := c.QueryParam("username")
	fmt.Println("Username:", username)


    tokenString := c.Request().Header.Get("Authorization")

    if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization header is missing or invalid"})
    }


    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    secretKey := []byte("your-secret-key")
    token, err := ParseToken(tokenString, secretKey)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse token"})
    }


fmt.Println("TOKEN")



fmt.Println(token)


return c.JSON(http.StatusOK, "userData")
}
*/

/*

func PersonalData(c echo.Context) error {
    username := c.QueryParam("username")
    fmt.Println("Username:", username)

    // Retrieve token string from Authorization header
    tokenString := c.Request().Header.Get("Authorization")
	fmt.Println(tokenString)
    if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization header is missing or invalid"})
    }
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    // Parse token
    secretKey := []byte("your-secret-key")
    token, err := ParseToken(tokenString, secretKey)
    if err != nil {
        // Handle token parsing error
        fmt.Println("Error parsing token:", err)
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse token"})
    }

    // Check if token is nil
    if token == nil {
        fmt.Println("Token is nil")
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
    }

    // Print token for debugging
    fmt.Println("TOKEN:", token)

    // Return some dummy data for now
    return c.JSON(http.StatusOK, map[string]string{"message": "Token parsed successfully"})
}
*/
/*
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

	// Extract the token string
	tokenString = tokenJSON.Token

	// Parse token
	secretKey := []byte("your-secret-key")
	token, err := ParseToken(tokenString, secretKey)
	if err != nil {
		// Handle token parsing error
		fmt.Println("Error parsing token:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse token"})
	}

	// Check if token is nil
	if token == nil {
		fmt.Println("Token is nil")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	// Print token for debugging
	fmt.Println("TOKEN:", token)

	usernameClaim, ok := token.Claims.(jwt.MapClaims)["username"]
	if !ok {
		fmt.Println("Username claim not found in token")
		// Handle the case where the username claim is missing
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Username claim not found in token"})
	}

	// Assert the username claim to string type
	username, ok = usernameClaim.(string)
	if !ok {
		fmt.Println("Invalid username claim in token")

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid username claim in token"})
	}


	fmt.Println("Username:", username)

	foundUsername, country, tel, token, chats, avatar, describtion, err := db.FindUserDataByUsername(username)
	fmt.Println(foundUsername, country, tel, token, chats, avatar, describtion, err )


	return c.JSON(http.StatusOK, map[string]string{"message": "Token parsed successfully"})
}
*/




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

	return c.JSON(http.StatusOK, map[string]string{"message": "Token parsed successfully"})
}

/*

	username := c.QueryParam("username")
	fmt.Println("Username:", username)

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username parameter is required"})
	}

	foundUsername, country, tel, token, chats, avatar, describtion, err := db.FindUserDataByUsername(username)
	fmt.Println(foundUsername, country, tel, token, chats, avatar, describtion, err )
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user data"})
	}




	userData := map[string]interface{}{
		"username":    foundUsername,
		"country":     country,
		"tel":         tel,
		"token":       token,
		"chats":       chats,
		"avatar":      avatar,
		"description": describtion,
	}
	return c.JSON(http.StatusOK, userData) */
