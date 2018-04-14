package chats

import (
	"db"
	"handlers"
	"views"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Create(ctx echo.Context) error {
	var chat views.Chat

	if err := ctx.Bind(&chat); err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err := ctx.Validate(&chat); err != nil {
		log.Infof("validation error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	current, ok := ctx.Get(handlers.CurrentUserKey).(db.User)
	if !ok {
		log.Info("can't get current user from context")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	usersMap := make(map[uint]*db.User)
	usersMap[current.ID] = &current

	for _, u := range chat.Users {
		if _, ok := usersMap[u.ID]; ok {
			continue
		}

		user := &db.User{
			ID: u.ID,
		}

		if err := db.Get(user); err != nil {
			if err == db.RecordNotFound {
				log.Infof("user with id %d not found", u.ID)
				return ctx.NoContent(http.StatusBadRequest)
			} else {
				log.Errorf("can't get current user from db: %v", err)
				return ctx.NoContent(http.StatusInternalServerError)
			}
		}

		usersMap[user.ID] = user
	}

	i := 0
	usersViews := make([]*views.User, len(usersMap))
	members := make([]*db.Member, len(usersMap))
	for _, u := range usersMap {
		usersViews[i] = views.NewUserView(u)
		members[i] = &db.Member{
			UserID: u.ID,
		}

		i++
	}

	model := &db.Chat{
		Name:    chat.Name,
		Members: members,
	}

	if err := db.Create(model); err != nil {
		log.Errorf("create chat error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := views.NewChatView(model, usersViews, nil)

	return handlers.JSONApiResponse(ctx, &view, http.StatusCreated)
}
