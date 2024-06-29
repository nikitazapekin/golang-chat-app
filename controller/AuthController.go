package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"
	"net/http"
	"strings"
	"time"
	//"log"
	"math/rand"
)

var currentAccessToken string 

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


func generateRandomColor() string {
    rand.Seed(time.Now().UnixNano())
    colors := []string{
        "orange", "orangered", "darkorange",
        "blue", "deepskyblue", "dodgerblue", 
        "green", "limegreen", "seagreen",   
    }
    return colors[rand.Intn(len(colors))]
}
/*
func CreateUserTable(username string) {
	fmt.Println("CTEATING NEW TABLE")
	if DB == nil {
		log.Fatal("Database connection is not established. Call Connect function first.")
	}
	tableName := fmt.Sprintf("user_data_%s", username)
	query := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        chats JSONB, 
        user_id VARCHAR(255)
    );
    `, tableName)
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Printf("Table %s created successfully.\n", tableName)
}
 */
func Register(c echo.Context) error {
	var registrationData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&registrationData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	username, country, tel, err := db.FindUserByUsername(registrationData.Username)

	fmt.Println("USEEEEEEEEER"+username, country, tel, err)
	if err != nil {
		if strings.Contains(err.Error(), "user with username") {
			db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
			fmt.Println("CREAATINGGGGGGGGGGGGGGGGG USERRRRRRRRRRRRRRRRRRRRRRr", registrationData.Username)



			fmt.Println(registrationData.Username)
	//		db.CreateUserTable(registrationData.Username)
			//fmt.Println(username)
			//db.CreateUserTable(username)
		} else {
			fmt.Println("Error finding user:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to find user"})
		}
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT access token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
		"exp": time.Now().Add(200 * time.Minute).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT refresh token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate refresh token"})
	}
	db.UpdateUserToken(registrationData.Username, refreshTokenString)
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenString,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Response().Writer, &cookie)
	response := RegistrationResponse{
		Username:     registrationData.Username,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return c.JSON(http.StatusOK, response)
}
func GetRefreshTokenExpirationTime(refreshToken string) (int64, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}
	if !token.Valid {
		return 0, fmt.Errorf("token is invalid")
	}
	expirationTime := token.Claims.(jwt.MapClaims)["exp"].(float64)
	return int64(expirationTime), nil
}
func isTokenValid(expirationTime int64) bool {
	currentTime := time.Now().Unix()
	if currentTime > expirationTime {
		return false
	}
	return true
}

func GetAccessToken(c echo.Context) error {
	token := c.QueryParam("token")
	user := c.QueryParam("user")
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Token not found in query"})
	}
	foundUsername, country, tel, refreshToken, chats, avatar, description, err := db.FindUserDataByUsername(user)
	fmt.Println(foundUsername, country, tel, refreshToken, chats, avatar, description, err)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка в поиске юзера"})
	}
	if foundUsername == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Пользователь не найден"})
	}
	expirationTime, err := GetRefreshTokenExpirationTime(refreshToken)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ошибка вытягивания времени из токена"})
	}
	if !isTokenValid(expirationTime) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Токен истек"})
	}
	newAccessToken, err := generateAccessToken(user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка генерации аксес токена"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": newAccessToken})
}
func tokenExpired(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return time.Now().After(expirationTime)
	} else {
		fmt.Println("Error parsing token:", token)
		return true 
	}
}
func generateAccessToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func RefreshAccessToken(c echo.Context) error {
	refreshTokenString := c.FormValue("refresh_token")
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})
	if err != nil || !refreshToken.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse refresh token claims"})
	}
	username := claims["username"].(string)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}
	response := map[string]string{
		"access_token": accessTokenString,
	}
	if err := db.UpdateAccessToken(username, accessTokenString); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update access token in the database"})
	}

	return c.JSON(http.StatusOK, response)
}
func GetRefreshToken(c echo.Context) error {
	cookie, err := c.Request().Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Refresh token not found"})
	}
	refreshToken := cookie.Value
	response := map[string]string{
		"refresh_token": refreshToken,
	}
	return c.JSON(http.StatusOK, response)
}
func GetAccessTokenStart(c echo.Context) error {
	user := c.QueryParam("user")
	foundUsername, country, tel, refreshToken, chats, avatar, description, err := db.FindUserDataByUsername(user)
	fmt.Println(foundUsername, country, tel, refreshToken, chats, avatar, description, err)
	return c.JSON(http.StatusOK, map[string]string{"token": "dfd"})
}
