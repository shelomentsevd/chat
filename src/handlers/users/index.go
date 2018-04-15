package users

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
	var (
		err   error
		users []*db.User
	)

	str := ctx.Param("chat")
	if str != "" {
		id, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			log.Infof("parse error: %v", err)
			return ctx.NoContent(http.StatusBadRequest)
		}

		users, err = chatUsers(uint(id), pagenator)
	} else {
		users, err = allUsers(pagenator)
	}

	if err != nil {
		log.Errorf("database error: %v", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	viewCollection := make([]*views.User, len(users))
	for i, u := range users {
		viewCollection[i] = views.NewUserView(u)
	}

	return handlers.JSONApiResponse(ctx, viewCollection, http.StatusOK)
}

func chatUsers(id uint, pagenator *pagenators.Pagenator) ([]*db.User, error) {
	var members []*db.Member

	if err := db.Get(&members,
		db.WithCondition(&db.Member{
			ChatID: id,
		}),
		db.WithLimit(pagenator.Limit),
		db.WithOffset(pagenator.Offset)); err != nil {
		return nil, err
	}

	ids := make([]uint, len(members))
	for i, m := range members {
		ids[i] = m.UserID
	}

	var users []*db.User
	if err := db.Get(&users, db.WithIDs(ids...)); err != nil {
		return nil, err
	}

	return users, nil
}

func allUsers(pagenator *pagenators.Pagenator) ([]*db.User, error) {
	var users []*db.User

	if err := db.Get(&users, db.WithLimit(pagenator.Limit), db.WithOffset(pagenator.Offset)); err != nil {
		return nil, err
	}

	return users, nil
}
