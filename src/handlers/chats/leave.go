package chats

import (
	"db"
	"handlers"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Leave(ctx echo.Context) error {
	user, ok := ctx.Get(handlers.CurrentUserKey).(db.User)
	if !ok {
		log.Info("can't add non-authorized user to chat")
		return ctx.NoContent(http.StatusUnauthorized)
	}

	chatStr := ctx.Param("chat")
	chatID, err := strconv.ParseUint(chatStr, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := db.Get(&db.Chat{
		ID: uint(chatID),
	}); err != nil {
		switch err {
		case db.ErrRecordNotFound:
			return ctx.NoContent(http.StatusBadRequest)
		default:
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	err = db.Delete(&db.Member{
		UserID: user.ID,
		ChatID: uint(chatID),
	})
	if err != nil {
		switch err {
		case db.ErrPgForeignKeyViolation:
			log.Infof("parse error: %v", err)
			return ctx.NoContent(http.StatusBadRequest)
		case db.ErrPgUniqueViolation:
		default:
			log.Errorf("database error: %v", err)
			return ctx.NoContent(http.StatusInternalServerError)
		}
	}

	return ctx.NoContent(http.StatusOK)
}
