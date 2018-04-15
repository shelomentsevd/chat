package users

import (
	"configuration"
	"db"
	"handlers"

	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func index_request() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/chats", nil)
}

func TestIndex(t *testing.T) {
	config, err := configuration.New()
	if err != nil {
		t.Errorf("tests error: %v", err)
	}

	if err := db.Init(config.DB.Params(),
		config.DB.MaxConnections,
		config.DB.MaxIdleConnections,
		config.DB.ConnectionLifeTime); err != nil {
		t.Errorf("to perform tests you should give database connection: %v", err)
	}

	pool := db.Pool

	Convey("Index users test", t, func() {
		db.Pool = pool.Begin()
		defer db.Pool.Rollback()

		e := echo.New()
		e.Validator = handlers.NewValidator()
		e.Binder = &handlers.Binder{
			Default: new(echo.DefaultBinder),
		}

		pwd := "Password"
		alice := db.User{
			Name:     "Alice",
			Password: &pwd,
		}
		if err := db.Create(&alice); err != nil {
			t.Error(err)
		}

		users := []*db.User{
			&db.User{Name: "User#1", Password: &pwd},
			&db.User{Name: "User#2", Password: &pwd},
			&db.User{Name: "User#3", Password: &pwd},
			&db.User{Name: "Bob", Password: &pwd},
		}

		for _, user := range users {
			if err := db.Create(user); err != nil {
				t.Error(err)
			}
		}

		chat := db.Chat{
			Name: "Chat one#001",
			Members: []*db.Member{
				&db.Member{
					UserID: alice.ID,
				},
				&db.Member{
					UserID: users[0].ID,
				},
			},
		}

		if err := db.Create(&chat); err != nil {
			t.Error(err)
		}

		Convey("All users", func() {
			req := index_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Index(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Convey("All users in chat", func() {
			req := index_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
			ctx.SetParamNames("chat")

			err := Index(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})
	})
}
