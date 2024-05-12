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
	// "github.com/example/jwtlib"
)

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
			// User not found, let's add them
			db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
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
		//	"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"exp": time.Now().Add(2 * time.Minute).Unix(),
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
	// Парсим токен, игнорируя проверку подписи
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	// Проверяем, что токен валиден
	if !token.Valid {
		return 0, fmt.Errorf("token is invalid")
	}

	// Получаем время истечения из расшифрованного токена
	expirationTime := token.Claims.(jwt.MapClaims)["exp"].(float64)

	// Возвращаем время истечения токена refreshToken
	return int64(expirationTime), nil
}

// Проверка годности токена
func isTokenValid(expirationTime int64) bool {
	// Получаем текущее время в Unix формате
	currentTime := time.Now().Unix()

	// Сравниваем время истечения токена с текущим временем
	if currentTime > expirationTime {
		// Если текущее время больше времени истечения токена, то токен просрочен
		return false
	}
	// Токен действителен
	return true
}

func GetAccessToken(c echo.Context) error {
	fmt.Println("TOKEEEEEEEEEEN")

	// Получаем токен из запроса
	token := c.QueryParam("token")
	user := c.QueryParam("user")
	fmt.Println("USSS" + user)
	fmt.Println(token)

	// Проверяем, был ли передан токен в запросе
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Token not found in query"})
	}

	// Проверяем валидность refresh токена и получаем данные пользователя
	foundUsername, country, tel, refreshToken, chats, avatar, description, err := db.FindUserDataByUsername(user)
	fmt.Println(foundUsername, country, tel, refreshToken, chats, avatar, description, err)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка в поиске юзера"})
	}

	// Проверяем, найден ли пользователь
	if foundUsername == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Пользователь не найден"})
	}

	// Получаем время истечения токена refreshToken
	expirationTime, err := GetRefreshTokenExpirationTime(refreshToken)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Ошибка вытягивания времени из токена"})
	}

	// Проверяем годность токена
	if !isTokenValid(expirationTime) {
		// Если токен просрочен, отправляем ошибку 401 Unauthorized
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Токен истек"})
	}

	newAccessToken, err := generateAccessToken(user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка генерации аксес токена"})
	}

	// Отправляем новый access токен на клиент
	fmt.Println("ACCEEESSSS", newAccessToken)
	//return c.JSON(http.StatusOK, map[string]string{"access_token": newAccessToken})
	return c.JSON(http.StatusOK, map[string]string{"token": newAccessToken})
}

func tokenExpired(tokenString string) bool {
	// Парсим токен для получения времени его истечения
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Предполагаем, что токен был подписан с использованием HMAC и имеет секретный ключ "my_secret_key"
		return []byte("your-secret-key"), nil
	})

	// Проверяем, была ли ошибка при парсинге токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Получаем время истечения токена из его утверждений
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		// Сравниваем время истечения токена с текущим временем
		return time.Now().After(expirationTime)
	} else {
		fmt.Println("Error parsing token:", token)
		return true // Если возникла ошибка при парсинге токена, считаем его истекшим
	}
}

// Функция для генерации нового access токена
func generateAccessToken(user string) (string, error) {
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем данные пользователя в токен
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user
	// Устанавливаем срок годности токена (1 час)
	//	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshAccessToken(c echo.Context) error {
	// Получаем refresh токен из запроса
	refreshTokenString := c.FormValue("refresh_token")

	// Проверяем валидность refresh токена
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil || !refreshToken.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
	}

	// Извлекаем из refresh токена данные, например, имя пользователя
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse refresh token claims"})
	}

	username := claims["username"].(string)

	// Генерируем новый access токен
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(), // Устанавливаем время истечения срока действия на 1 минуту
	})
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}

	// Возвращаем новый access токен клиенту
	response := map[string]string{
		"access_token": accessTokenString,
	}

	return c.JSON(http.StatusOK, response)
}

func GetRefreshToken(c echo.Context) error {
	// Получаем значение refresh_token из http-only cookie
	cookie, err := c.Request().Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Refresh token not found"})
	}

	refreshToken := cookie.Value

	// Возвращаем значение refresh_token клиенту
	response := map[string]string{
		"refresh_token": refreshToken,
	}

	return c.JSON(http.StatusOK, response)
}

func GetAccessTokenStart(c echo.Context) error {
	fmt.Println("TOKEEEEEEEEEEN")
	user := c.QueryParam("user")
	fmt.Println("USSsssS" + user)
	foundUsername, country, tel, refreshToken, chats, avatar, description, err := db.FindUserDataByUsername(user)
	fmt.Println(foundUsername, country, tel, refreshToken, chats, avatar, description, err)
	return c.JSON(http.StatusOK, map[string]string{"token": "dfd"})
}
