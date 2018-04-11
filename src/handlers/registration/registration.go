package registration

import (
	"db"
	"db/users"

	"bytes"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func RegisterUser(ctx echo.Context) error {
	var user db.User
	if err := ctx.Bind(&user); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(&user); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := users.Create(&user); err != nil {
		log.Infof("create user error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &user); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
