package pagenators

import (
	"strconv"

	"github.com/labstack/echo"
)

const (
	defaultLimit = 25

	pageLimit  = "page[limit]"
	pageOffset = "page[offset]"
)

// TODO: Enhancement. Generate links like in example on http://jsonapi.org/ page.
type Pagenator struct {
	Limit  int
	Offset int
}

func NewPaginator(ctx echo.Context) *Pagenator {
	paramLimit := ctx.QueryParam(pageLimit)
	paramOffset := ctx.QueryParam(pageOffset)

	limit, _ := strconv.Atoi(paramLimit)
	offset, _ := strconv.Atoi(paramOffset)

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 || limit > defaultLimit {
		limit = defaultLimit
	}

	return &Pagenator{
		Offset: offset,
		Limit:  limit,
	}
}
