package messages

import (
	"db"
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

	err = db.Get(chatModel)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return ctx.NoContent(http.StatusBadRequest)
		} else {
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	current, ok := ctx.Get(handlers.CurrentUserKey).(db.User)
	if !ok {
		log.Error("can't create message by unauthorized user")
		return ctx.NoContent(http.StatusUnauthorized)
	}

	var members []*db.Member
	err = db.Get(&members, db.WithCondition(&db.Member{
		ChatID: chatModel.ID,
		UserID: current.ID,
	}))
	if len(members) == 0 {
		return ctx.NoContent(http.StatusForbidden)
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

	messageModel := &db.Message{
		Content:   newMessage.Content,
		ChatID:    chatModel.ID,
		UserID:    current.ID,
		User:      &current,
		CreatedAt: time.Now(),
	}

	if err := db.Create(messageModel); err != nil {
		log.Errorf("can't create message error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := views.NewMessageView(messageModel, views.NewUserView(messageModel.User))

	return handlers.JSONApiResponse(ctx, view, http.StatusCreated)
}
