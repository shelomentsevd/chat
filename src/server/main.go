package main

import (
	"configuration"
	"db"
	"handlers"
	"handlers/authorization"
	"handlers/chats"
	"handlers/messages"
	"handlers/registration"

	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tylerb/graceful"
)

func main() {
	config, err := configuration.New()
	if err != nil {
		log.Fatalf("can't read configuration error: %v", err)
	}

	if err := db.Init(config.DB.Params(), config.DB.MaxConnections, config.DB.MaxIdleConnections, config.DB.ConnectionLifeTime); err != nil {
		log.Fatalf("can't connect to database by URL: %s error: %v", config.DB.Params(), err)
	}

	e := echo.New()
	e.Binder = &handlers.Binder{
		Default: new(echo.DefaultBinder),
	}

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

	e.Server.Addr = config.Server.String()

	graceful.ListenAndServe(e.Server, 10*time.Second)
}
