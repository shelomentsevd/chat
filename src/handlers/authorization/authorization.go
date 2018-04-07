package authorization

import (
	"db/users"
	"handlers"

	"fmt"

	"github.com/labstack/echo"
)

func BasicAuthValidator(user, password string, c echo.Context) (bool, error) {
	ctx, ok := c.(handlers.Context)

	if !ok {
		return false, fmt.Errorf("can't cast %T to %T", c, handlers.Context{})
	}

	model, err := users.GetByName(user)
	if err != nil {
		return false, err
	}

	if model.Password != password {
		return false, nil
	}

	ctx.User = model

	return true, nil
}
