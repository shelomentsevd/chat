package authorization

import (
	"db/users"
	"models"

	"github.com/labstack/echo"
)

func BasicAuthValidator(username, password string, ctx echo.Context) (bool, error) {
	user := models.User{
		Name: username,
	}

	if err := users.Get(&user); err != nil {
		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	ctx.Set("current_user", user)

	return true, nil
}
