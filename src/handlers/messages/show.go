package messages

import (
	"db"
	"handlers"
	"views"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Show(ctx echo.Context) error {
	chatStr := ctx.Param("chat")
	chatID, err := strconv.ParseUint(chatStr, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	messageStr := ctx.Param("message")
	messageID, err := strconv.ParseUint(messageStr, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	model := &db.Message{
		ID:     uint(messageID),
		ChatID: uint(chatID),
	}

	err = db.Get(model)
	switch err {
	case db.ErrRecordNotFound:
		return ctx.NoContent(http.StatusNotFound)
	default:
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	userModel := &db.User{
		ID: model.UserID,
	}

	err = db.Get(userModel)
	if err != nil {
		log.Errorf("can't find message author database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := views.NewMessageView(model, views.NewUserView(userModel))

	return handlers.JSONApiResponse(ctx, &view, http.StatusOK)
}
