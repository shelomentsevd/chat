package handlers

import "github.com/labstack/echo"

type Binder struct{}

func (b *Binder) Bind(i interface{}, c echo.Context) (err error) {

	return nil
}
