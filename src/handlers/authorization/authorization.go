package authorization

import (
	"handlers"

	"fmt"

	"github.com/labstack/echo"
)

func BasicAuthValidator(user, password string, c echo.Context) (bool, error) {
	// TODO: Прикрутить авторизацию
	_, ok := c.(handlers.Context)

	if !ok {
		return false, fmt.Errorf("can't cast %T to %T", c, handlers.Context{})
	}

	return true, nil
}