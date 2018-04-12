package chats

import (
	"db"
	"db/chats"
	"handlers"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const chatLimit = 25

func Index(ctx echo.Context) error {
	var query indexParameters

	if err := ctx.Bind(&query); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if query.Limit < 0 || query.Limit > chatLimit {
		query.Limit = chatLimit
	}

	if query.Offset < 0 {
		query.Offset = 0
	}

	var slice []*db.Chat

	if err := chats.Get(slice, query.Limit, query.Offset); err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return handlers.JSONApiResponse(ctx, &slice, http.StatusOK)
}
