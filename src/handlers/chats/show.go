package chats

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
	str := ctx.Param("chat")

	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	model := &db.Chat{
		ID: uint(id),
	}

	err = db.Get(model)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return ctx.NoContent(http.StatusNotFound)
		} else {
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	view := views.NewChatView(model, nil, nil)

	return handlers.JSONApiResponse(ctx, &view, http.StatusOK)
}
