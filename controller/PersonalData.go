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
 









/*

package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"

	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	//"strings"
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

	// Retrieve refresh token from the database
	foundUsername, country, tel, refreshToken, chats, avatar, describtion, err := db.FindUserDataByUsername(username)
	if err != nil {
		fmt.Println("Error fetching user data:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user data"})
	}

	// Check if refresh token is empty
	if refreshToken == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Refresh token not found"})
	}
fmt.Println("REFRESH")
fmt.Println(refreshToken)
	// Parse the JSON to extract the token string
	var tokenJSON struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal([]byte(refreshToken), &tokenJSON); err != nil {
		fmt.Println("Error decoding token JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token format"})
	}

	// Extract the token string
	refreshToken = tokenJSON.Token

	// Parse and validate the refresh token
	secretKey := []byte("your-secret-key")
	token, err := ParseToken(refreshToken, secretKey)
	if err != nil {
		fmt.Println("Error parsing refresh token:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse refresh token"})
	}

	// Check if the token is nil or expired
	if token == nil {
		fmt.Println("Refresh token is nil")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
	}

	// Get the expiration time of the token
	expirationTime := token.Claims.(jwt.MapClaims)["exp"].(float64)
	expiration := time.Unix(int64(expirationTime), 0)

	// Check if the token has expired
	if time.Now().After(expiration) {
		fmt.Println("Refresh token has expired")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Refresh token has expired"})
	}

	// Fetch additional user data or process existing data
	// For example:
	// avatar, description, err := db.FindAvatarAndDescriptionByUsername(username)
	// if err != nil {
	// 	fmt.Println("Error fetching avatar and description:", err)
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch avatar and description"})
	// }

	// Form the response data
	responseData := map[string]interface{}{
		"access_token": "newAccessToken",
		"username":     foundUsername,
		"country":      country,
		"tel":          tel,
		"chats":        chats,
		"avatar":       avatar,
		"description":  describtion,
	}

	// Return the user data
	return c.JSON(http.StatusOK, responseData)
}


*/
















/*
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

// Function to parse a token
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

// Function to generate a new access token
func GenerateAccessToken(username string, secretKey []byte) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    
    accessToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    
    return accessToken, nil
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

    // Parse and validate the refresh token
    refreshSecretKey := []byte("your-secret-key")
    refreshToken, err := ParseToken(tokenString, refreshSecretKey)
    if err != nil {
        fmt.Println("Error parsing refresh token:", err)
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to parse refresh token"})
    }

    if refreshToken == nil {
        fmt.Println("Refresh token is nil")
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
    }

    // Extract claims from the refresh token
    refreshUsernameClaim, ok := refreshToken.Claims.(jwt.MapClaims)["username"]
    if !ok {
        fmt.Println("Username claim not found in refresh token")
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Username claim not found in refresh token"})
    }

    refreshUsername, ok := refreshUsernameClaim.(string)
    if !ok {
        fmt.Println("Invalid username claim in refresh token")
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid username claim in refresh token"})
    }

    // Generate a new access token
    newAccessToken, err := GenerateAccessToken(refreshUsername, secretKey)
    if err != nil {
        fmt.Println("Error generating access token:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
    }

    // Retrieve user data using the username from the refresh token
    // Assuming db.FindUserDataByUsername returns foundUsername, country, tel, chats, avatar, description, and err
    foundUsername, country, tel, _, chats, avatar, description, err := db.FindUserDataByUsername(refreshUsername)
    if err != nil {
        fmt.Println("Error retrieving user data:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user data"})
    }

    // Return the new access token and user data to the client
    responseData := map[string]interface{}{
        "access_token": newAccessToken,
        "username":     foundUsername,
        "country":      country,
        "tel":          tel,
        "chats":        chats,
        "avatar":       avatar,
        "description":  description,
    }

    return c.JSON(http.StatusOK, responseData)
}
*/