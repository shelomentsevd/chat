package messages

import (
	"configuration"
	"db"
	"handlers"

	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func show_request() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/chats", nil)
}

func TestShow(t *testing.T) {
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

	Convey("Show message test", t, func() {
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

		messages := []*db.Message{
			&db.Message{
				Content:   "Message content",
				ChatID:    chat.ID,
				UserID:    users[0].ID,
				CreatedAt: time.Now(),
			},
			&db.Message{
				Content:   "Message content 2",
				ChatID:    chat.ID,
				UserID:    users[1].ID,
				CreatedAt: time.Now(),
			},
			&db.Message{
				Content:   "Message content 3",
				ChatID:    chat.ID,
				UserID:    users[2].ID,
				CreatedAt: time.Now(),
			},
			&db.Message{
				Content:   "Message content 4",
				ChatID:    chat.ID,
				UserID:    users[3].ID,
				CreatedAt: time.Now(),
			},
			&db.Message{
				Content:   "Message content 5",
				ChatID:    chat.ID,
				UserID:    users[1].ID,
				CreatedAt: time.Now(),
			},
		}

		for _, m := range messages {
			if err := db.Create(m); err != nil {
				t.Error(err)
			}
		}

		Convey("Message exists", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10), strconv.FormatUint(uint64(messages[0].ID), 10))
			ctx.SetParamNames("chat", "message")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Convey("Message doesn't exist", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10), "32987")
			ctx.SetParamNames("chat", "message")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("Chat doesn't exist", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues("9823", strconv.FormatUint(uint64(messages[0].ID), 10))
			ctx.SetParamNames("chat", "message")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusNotFound)
		})
	})
}
