package chats

import (
	"db/chats"
	"models"

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

	chat := models.Chat{
		ID: uint(id),
	}

	// TODO: ErrNotFound processing!
	if err := chats.GetByID(&chat); err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &chat); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
