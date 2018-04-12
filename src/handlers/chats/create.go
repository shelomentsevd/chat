package chats

import (
	"db"
	"db/chats"
	"handlers"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Create(ctx echo.Context) error {
	var chat db.Chat

	if err := ctx.Bind(&chat); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(&chat); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := chats.Create(&chat); err != nil {
		log.Infof("create chat error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	return handlers.JSONApiResponse(ctx, &chat, http.StatusCreated)
}
