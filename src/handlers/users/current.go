package users

import (
	"db"
	"handlers"
	"views"

	"net/http"

	"github.com/labstack/echo"
)

func Current(ctx echo.Context) error {
	user := ctx.Get(handlers.CurrentUserKey).(db.User)

	view := views.NewUserView(&user)

	return handlers.JSONApiResponse(ctx, view, http.StatusOK)
}
