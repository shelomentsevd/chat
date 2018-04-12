package chats

import (
	"db"
	"db/chats"
	"handlers"
	"views"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Create(ctx echo.Context) error {
	var chat views.Chat

	if err := ctx.Bind(&chat); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(&chat); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	// TODO: Check users, and add them
	members := make([]*db.Member, len(chat.Users))
	for i, u := range chat.Users {
		members[i] = &db.Member{
			UserID: u.ID,
		}
	}

	model := &db.Chat{
		Name:    chat.Name,
		Members: members,
	}

	if err := chats.Create(model); err != nil {
		log.Infof("create chat error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	view := views.NewChatView(model)

	return handlers.JSONApiResponse(ctx, &view, http.StatusCreated)
}
