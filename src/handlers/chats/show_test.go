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

	db.Pool = db.Pool.Begin()
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
		&db.User{Name: "User#4", Password: &pwd},
		&db.User{Name: "User#5", Password: &pwd},
		&db.User{Name: "User#6", Password: &pwd},
		&db.User{Name: "User#7", Password: &pwd},
		&db.User{Name: "User#8", Password: &pwd},
		&db.User{Name: "User#9", Password: &pwd},
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

	Convey("Chat show tests", t, func() {
		Convey("Chat exists", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
			ctx.SetParamNames("chat")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Convey("Chat does not exist", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamValues("1000")
			ctx.SetParamNames("chat")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("ID is not a number", func() {
			req := show_request()
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamValues("whatever")
			ctx.SetParamNames("chat")

			err := Show(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
