package chats

import (
	"db/chats"
	"models"

	"bytes"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Create(ctx echo.Context) error {
	var chat models.Chat

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

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &chat); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
