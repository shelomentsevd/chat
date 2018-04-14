package chats

import (
	"configuration"
	"db"
	"handlers"
	"views"

	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func request(json interface{}) *http.Request {
	out := bytes.NewBuffer(nil)
	jsonapi.MarshalPayload(out, json)

	r := httptest.NewRequest(http.MethodPost, "/chats", out)
	r.Header.Add("Content-Type", jsonapi.MediaType)

	return r
}

func TestCreate(t *testing.T) {
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

	bob := db.User{
		Name:     "Bob",
		Password: &pwd,
	}
	if err := db.Create(&bob); err != nil {
		t.Error(err)
	}

	Convey("Chat create tests", t, func() {
		Convey("Create chat with Bob", func() {
			chat := views.Chat{
				Name: "Alice & Bob chat",
				Users: []*views.User{
					views.NewUserView(&bob),
				},
			}

			req := request(&chat)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusCreated)
		})

		Convey("Create chat with himself and Bob", func() {
			chat := views.Chat{
				Name: "Alice & Bob chat 2",
				Users: []*views.User{
					views.NewUserView(&bob),
				},
			}

			req := request(&chat)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusCreated)
		})

		Convey("Create chat with himself only", func() {
			chat := views.Chat{
				Name:  "Alice independence chat",
				Users: nil,
			}

			req := request(&chat)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("Create chat with non-existing user", func() {
			chat := views.Chat{
				Name: "Alice independence chat",
				Users: []*views.User{
					&views.User{
						ID: 1000,
					},
				},
			}

			req := request(&chat)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("Create chat without title", func() {
			chat := views.Chat{
				Name: "",
				Users: []*views.User{
					views.NewUserView(&bob),
				},
			}

			req := request(&chat)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
