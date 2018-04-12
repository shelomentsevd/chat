package handlers

import (
	"bytes"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func JSONApiResponse(ctx echo.Context, response interface{}, status int) error {
	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, response); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(status, jsonapi.MediaType, out.Bytes())
}
