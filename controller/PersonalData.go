
package controller

import (
	//"encoding/json"
	"fmt"
	"net/http"
	//"time"

	"github.com/labstack/echo/v4"
	//"github.com/nikita/go-microservices/db"
	//"github.com/dgrijalva/jwt-go"
)


/*
var currentAccessToken string // Глобальная переменная для хранения текущего access токена

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type RegistrationResponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegistrationParams struct {
	Username  string `json:"username"`
	Country   string `json:"country"`
	Telephone string `json:"tel"`
}
*/
func PersonalData(c echo.Context) error {
	fmt.Println("personal")
/*	var registrationData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&registrationData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	fmt.Println("Received registration data:", registrationData)
	db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
//	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"username": registrationData.Username,
//	})


accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": registrationData.Username,
    "exp":      time.Now().Add(1 * time.Minute).Unix(), // Устанавливаем время истечения срока действия на 1 минуту
})
accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))

	
	currentAccessToken = accessTokenString

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT refresh token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate refresh token"})
	}

	// Set refresh token as HTTP-only cookie
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenString,
		Path:     "/",
		HttpOnly: true,
	//	Expires:  time.Now().Add(24 * time.Hour), // Set expiry time as per your requirement
	}
	http.SetCookie(c.Response().Writer, &cookie)

	response := RegistrationResponse{
		Username:    registrationData.Username,
		AccessToken: accessTokenString,
		//Expires: time.Now().Add(1 * time.Minute),
			RefreshToken: refreshTokenString,
	}
	return c.JSON(http.StatusOK, response)*/
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})


}
