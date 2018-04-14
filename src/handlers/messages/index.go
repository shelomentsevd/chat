package messages

import (
	"db"
	"handlers"
	"pagenators"
	"views"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Index(ctx echo.Context) error {
	pagenator := pagenators.NewPaginator(ctx)

	str := ctx.Param("chat")

	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Infof("parse error: %v", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	messagesModels := make([]*db.Message, 0)

	err = db.Get(messagesModels,
		db.WithCondition(&db.Message{
			ChatID: uint(id),
		}),
		db.WithLimit(pagenator.Limit),
		db.WithOffset(pagenator.Offset),
		db.WithOrder("created_at desc"))

	switch err {
	case db.RecordNotFound:
		log.Infof("chat %d not found", id)
		return ctx.NoContent(http.StatusNotFound)
	default:
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	ids := make([]uint, 0)
	usersMap := make(map[uint][]*db.Message)
	for _, m := range messagesModels {
		if userMessages, ok := usersMap[m.UserID]; ok {
			userMessages = append(userMessages, m)
			continue
		}

		usersMap[m.UserID] = make([]*db.Message, 1)
		usersMap[m.UserID][0] = m
		ids = append(ids, m.UserID)
	}

	var usersModels []*db.User
	if err := db.Get(usersModels, db.WithIDs(ids...)); err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	for _, u := range usersModels {
		userMessages := usersMap[u.ID]
		for _, m := range userMessages {
			m.User = u
		}
	}

	messagesViews := make([]*views.Message, len(messagesModels))
	for i, m := range messagesModels {
		messagesViews[i] = views.NewMessageView(m, views.NewUserView(m.User))
	}

	return handlers.JSONApiResponse(ctx, &messagesViews, http.StatusOK)
}
