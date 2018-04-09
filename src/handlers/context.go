package handlers

import (
	"models"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context

	// TODO: You can use Get() interface{} instead!
	// Current user
	User *models.User
}
