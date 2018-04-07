package main

import (
	"handlers/authorization"
	"handlers/chats"
	"handlers/messages"
	"handlers/registration"

	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Group api
	api := e.Group("/api/v1")
	// Registration
	api.POST("/registration", registration.RegisterUser)

	// Group chats
	chatsGroup := api.Group("/chats")
	chatsGroup.Use(middleware.BasicAuth(authorization.BasicAuthValidator))
	// chats routes
	chatsGroup.GET("/", chats.Index)
	chatsGroup.POST("/", chats.Create)
	chatsGroup.GET("/:chat", chats.Show)
	chatsGroup.GET("/:chat/join", chats.Join)
	chatsGroup.GET("/:chat/leave", chats.Leave)
	// message routes
	chatsGroup.GET("/:chat/messages", messages.Index)
	chatsGroup.GET("/:chat/messages/:message", messages.Show)
	chatsGroup.POST("/:chat/messages", messages.Create)

	e.Server.Addr = ":3000"

	graceful.ListenAndServe(e.Server, 10*time.Second)
}
