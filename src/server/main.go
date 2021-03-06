package main

import (
	"configuration"
	"db"
	"handlers"
	"handlers/chats"
	"handlers/messages"
	"handlers/registration"
	"handlers/users"

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

	// TODO: Refactor this
	for i := 0; i < 100; i++ {
		err := db.Init(config.DB.Params(), config.DB.MaxConnections, config.DB.MaxIdleConnections, config.DB.ConnectionLifeTime)
		if err == nil {
			break
		}
		log.Infof("error: %v", err)
		<-time.After(time.Second)
	}
	if err != nil {
		log.Fatalf("can't connect to database by URL: %s error: %v", config.DB.Params(), err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Validator = handlers.NewValidator()

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
	// Group users
	usersGroup := api.Group("/users")
	usersGroup.Use(middleware.BasicAuth(handlers.BasicAuthValidator))
	usersGroup.GET("/", users.Index)
	usersGroup.GET("/current", users.Current)
	// Group chats
	chatsGroup := api.Group("/chats")
	chatsGroup.Use(middleware.BasicAuth(handlers.BasicAuthValidator))
	// chats routes
	chatsGroup.GET("/", chats.Index)
	chatsGroup.POST("/", chats.Create)
	chatsGroup.GET("/:chat", chats.Show)
	chatsGroup.POST("/:chat/join", chats.Join)
	chatsGroup.POST("/:chat/leave", chats.Leave)
	chatsGroup.GET("/:chat/users", users.Index)
	// message routes
	chatsGroup.GET("/:chat/messages", messages.Index)
	chatsGroup.GET("/:chat/messages/:message", messages.Show)
	chatsGroup.POST("/:chat/messages", messages.Create)

	e.Server.Addr = config.Server.String()

	if err := graceful.ListenAndServe(e.Server, 10*time.Second); err != nil {
		log.Infof("Error after shutdown: %v", err)
	}
}
