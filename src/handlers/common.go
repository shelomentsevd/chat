package handlers

import (
	"db"

	"bytes"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const CurrentUserKey = "current_user"

func BasicAuthValidator(username, password string, ctx echo.Context) (bool, error) {
	user := db.User{
		Name: username,
	}

	if err := db.Get(&user); err != nil {
		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	ctx.Set(CurrentUserKey, user)

	return true, nil
}

func JSONApiResponse(ctx echo.Context, response interface{}, status int) error {
	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, response); err != nil {
		log.Errorf("marshal error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Blob(status, jsonapi.MediaType, out.Bytes())
}
