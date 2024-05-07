/*
package controller
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/nikita/go-microservices/db"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
	fmt.Println("Received registration data:", registrationData)
	db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)

	 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})

 
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	fmt.Println("TOKEN" +tokenString)
	if err != nil {
		fmt.Println("Error signing JWT token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

 
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}
*/


/*
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type RegistrationResponse struct {
	Username string `json:"username"`
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
	fmt.Println("Received registration data:", registrationData)
	db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT access token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT refresh token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate refresh token"})
	}

	response := RegistrationResponse{
		Username: registrationData.Username,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return c.JSON(http.StatusOK, response)
}
*/

/*

package controller
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	db "github.com/nikita/go-microservices/db"
)
type RegistrationParams struct {
	Username string `json:"username"`
	Country  string `json:"country"`
	Telephone string `json:"tel"`
}
func Register(c echo.Context) error {
	var registrationData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&registrationData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	fmt.Println("Received registration data:", registrationData)
	db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
 
	return c.JSON(http.StatusOK, "{message: gggg}")
}
 */







/*
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"
	"github.com/dgrijalva/jwt-go"
)

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
	fmt.Println("Received registration data:", registrationData)
	db.AddUser(registrationData.Username, registrationData.Country, registrationData.Telephone)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT access token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": registrationData.Username,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error signing JWT refresh token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate refresh token"})
	}

	// Set refresh token as HTTP-only cookie
	fmt.Println("REFRESH TOKEN" +refreshTokenString)

	fmt.Println("ACCESS TOKEN" +accessTokenString)
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenString,
		Path:     "/",
		HttpOnly: true,
		Expires: time.Now().Add(1 * time.Minute),

		//Expires:  time.Now().Add(24 * time.Hour), // Set expiry time as per your requirement
	}
	http.SetCookie(c.Response().Writer, &cookie)

	response := RegistrationResponse{
		Username:     registrationData.Username,
		AccessToken:  accessTokenString,
	//	RefreshToken: refreshTokenString,
	}
	return c.JSON(http.StatusOK, response)
}
*/





package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nikita/go-microservices/db"
	"github.com/dgrijalva/jwt-go"
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
	return c.JSON(http.StatusOK, response)
}

// Обработчик GET запроса для получения текущего access токена


/*
func GetAccessToken(c echo.Context) error {
	// Проверяем, есть ли текущий access токен
	if currentAccessToken == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Access token not found"})
	}

	// Возвращаем текущий access токен клиенту
	response := map[string]string{
		"access_token": currentAccessToken,
	}

	return c.JSON(http.StatusOK, response)
}

*/














func GetAccessToken(c echo.Context) error {
    // Проверяем, есть ли текущий access токен
    if currentAccessToken == "" {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Access token not found"})
    }

    // Проверяем время действия access токена
    token, err := jwt.Parse(currentAccessToken, func(token *jwt.Token) (interface{}, error) {
        // Проверяем алгоритм подписи токена
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("your-secret-key"), nil
    })

    if err != nil || !token.Valid {
        // Если access токен недействителен или истек его срок действия,
        // то обновляем его с помощью функции RefreshAccessToken
        refreshedToken, err := RefreshAccessTokenInternal()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to refresh access token"})
        }
        currentAccessToken = refreshedToken
    }

    // Возвращаем текущий access токен клиенту
    response := map[string]string{
        "access_token": currentAccessToken,
    }

    return c.JSON(http.StatusOK, response)
}

func RefreshAccessTokenInternal() (string, error) {
    // Генерируем новый access токен
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": "example_user",
        "exp":      time.Now().Add(1 * time.Minute).Unix(), // Устанавливаем время истечения срока действия на 1 минуту
    })
    accessTokenString, err := accessToken.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }
    return accessTokenString, nil
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
