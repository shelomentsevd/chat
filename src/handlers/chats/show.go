package chats

import (
	"db"
	"db/chats"

	"bytes"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
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

	chat := db.Chat{
		ID: uint(id),
	}

	err = chats.GetByID(&chat)
	if err != nil {
		if err == db.RecordNotFound {
			return ctx.NoContent(http.StatusNotFound)
		} else {
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &chat); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
