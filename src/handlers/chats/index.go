package chats

import (
	"db"
	"db/chats"

	"bytes"
	"net/http"

	"github.com/google/jsonapi"
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

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &slice); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
