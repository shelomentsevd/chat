package messages

import (
	"db"
	"db/chats"
	"db/messages"
	"handlers"
	"views"

	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Create(ctx echo.Context) error {
	str := ctx.Param("chat")

	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	chatModel := &db.Chat{
		ID: uint(id),
	}

	err = chats.Get(chatModel, false)
	if err != nil {
		if err == db.RecordNotFound {
			return ctx.NoContent(http.StatusNotFound)
		} else {
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	var newMessage views.Message
	if err := ctx.Bind(&newMessage); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(newMessage); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	current, ok := ctx.Get("current_user").(db.User)
	if !ok {
		log.Error("current user isn't specified")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	messageModel := &db.Message{
		Content:   newMessage.Content,
		ChatID:    chatModel.ID,
		UserID:    current.ID,
		User:      &current,
		CreatedAt: time.Now(),
	}

	if err := messages.Create(messageModel); err != nil {
		log.Errorf("can't create message error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := views.NewMessageView(messageModel, views.NewUserView(messageModel.User))

	return handlers.JSONApiResponse(ctx, &view, http.StatusCreated)
}
