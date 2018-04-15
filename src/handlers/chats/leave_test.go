package chats

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

func leave_request() *http.Request {
	return httptest.NewRequest(http.MethodPost, "/chats", nil)
}

func TestLeave(t *testing.T) {
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

	Convey("Chat join tests", t, func() {
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
				&db.Member{
					UserID: users[1].ID,
				},
				&db.Member{
					UserID: users[2].ID,
				},
			},
		}

		if err := db.Create(&chat); err != nil {
			t.Error(err)
		}

		Convey("chat exists", func() {
			req := leave_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
			ctx.SetParamNames("chat")

			err := Leave(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)

			Convey("leave again", func() {
				req := leave_request()
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.Set(handlers.CurrentUserKey, alice)
				ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
				ctx.SetParamNames("chat")

				err := Leave(ctx)
				So(err, ShouldBeNil)
				So(rec.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("Chat does not exist", func() {
			req := leave_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues("1000")
			ctx.SetParamNames("chat")

			err := Leave(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
