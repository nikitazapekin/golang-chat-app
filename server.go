//module github.com/nikita/go-microservices

package main

import (
	//"fmt"
	"log"
	"net/http"
	//"server/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/nikita/go-microservices/db"
	r "github.com/nikita/go-microservices/routes"
)

func main() {
	e := echo.New()
	db.Connect()
	//	e.Use(middleware.CORS())
	/*
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	*/
/*	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	*/
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	

	e.GET("/chat/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r.InitRoutes(e)
	/*
	 */
	r.InitWebsocketRoutes(e)
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

// go run server.go
// go mod init github.com/nikita/go-microservices
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get -u github.com/golang/protobuf/protoc-gen
//go get -u google.golang.org/grpc
// mkdir chat
//go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
