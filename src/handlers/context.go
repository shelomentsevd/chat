package handlers

import (
	"models"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context

	// Current user
	User *models.User
}
