package handlers

import (
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

type Binder struct {
	Default *echo.DefaultBinder
}

func (b *Binder) Bind(i interface{}, c echo.Context) error {
	req := c.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if ctype != jsonapi.MediaType {
		if err := b.Default.Bind(i, c); err != echo.ErrUnsupportedMediaType {
			return err
		}
	}

	if err := jsonapi.UnmarshalPayload(req.Body, i); err != nil {
		return err
	}

	return nil
}
