package chats

import (
	"db"
	"db/chats"
	"handlers"
	"pagenators"
	"views"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Index(ctx echo.Context) error {
	pagenator := pagenators.NewPaginator(ctx)

	var slice []*db.Chat

	if err := chats.GetList(slice, pagenator.Limit, pagenator.Offset); err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	viewCollection := make([]*views.Chat, len(slice))
	for i, c := range slice {
		viewCollection[i] = views.NewChatView(c, nil, nil)
	}

	return handlers.JSONApiResponse(ctx, &viewCollection, http.StatusOK)
}
