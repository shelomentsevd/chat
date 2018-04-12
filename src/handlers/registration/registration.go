package registration

import (
	"db"
	"db/users"
	"handlers"
	"views"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type form struct {
	Name     string `form:"name"     validate:"required"`
	Password string `form:"password" validate:"required"`
}

func RegisterUser(ctx echo.Context) error {
	var user form
	if err := ctx.Bind(&user); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(&user); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	model := &db.User{
		Name:     user.Name,
		Password: user.Password,
	}

	if err := users.Create(model); err != nil {
		log.Infof("create user error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	view := views.NewUserView(model)

	return handlers.JSONApiResponse(ctx, view, http.StatusCreated)
}
