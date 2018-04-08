package registration

import (
	"db/users"
	"models"

	"bytes"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Errorf("parse error: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := c.Validate(&user); err != nil {
		log.Infof("validation error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := users.CreateUser(&user); err != nil {
		log.Infof("create user error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	out := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalPayload(out, &user); err != nil {
		log.Errorf("marshal error: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Blob(http.StatusCreated, jsonapi.MediaType, out.Bytes())
}
