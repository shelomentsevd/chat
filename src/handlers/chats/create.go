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
		return ctx.NoContent(http.StatusUnauthorized)
	}

	ids := make([]uint, len(chat.Users))
	for i, u := range chat.Users {
		ids[i] = u.ID
	}

	var userModels []*db.User
	if err := db.Get(&userModels, db.WithIDs(ids...)); err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if len(userModels) == 0 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	usersMap := make(map[uint]*db.User)
	usersMap[current.ID] = &current
	for _, u := range userModels {
		usersMap[u.ID] = u
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

	view := views.NewChatView(model, usersViews)

	return handlers.JSONApiResponse(ctx, view, http.StatusCreated)
}
