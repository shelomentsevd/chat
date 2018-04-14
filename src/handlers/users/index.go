package users

import (
	"db"
	"db/users"
	"handlers"
	"pagenators"
	"views"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Index(ctx echo.Context) error {
	pagenator := pagenators.NewPaginator(ctx)

	var models []*db.User

	if err := users.GetList(models, pagenator.Limit, pagenator.Offset); err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	viewCollection := make([]*views.User, len(models))
	for i, u := range models {
		viewCollection[i] = views.NewUserView(u)
	}

	return handlers.JSONApiResponse(ctx, &viewCollection, http.StatusOK)
}
