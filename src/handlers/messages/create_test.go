package messages

import (
	"configuration"
	"db"
	"handlers"
	"views"

	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func create_request(json interface{}) *http.Request {
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

	pool := db.Pool

	Convey("Create message test", t, func() {
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

		Convey("Chat exists. User is a member.", func() {
			message := views.NewMessageView(&db.Message{
				Content: "Test message",
			}, nil)
			req := create_request(message)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
			ctx.SetParamNames("chat")

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusCreated)
		})

		Convey("Chat exists. User isn't a member.", func() {
			message := views.NewMessageView(&db.Message{
				Content: "Test message dsfdsfsdf",
			}, nil)
			req := create_request(message)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, *users[3])
			ctx.SetParamValues(strconv.FormatUint(uint64(chat.ID), 10))
			ctx.SetParamNames("chat")

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusForbidden)
		})

		Convey("Chat does not exist.", func() {
			message := views.NewMessageView(&db.Message{
				Content: "Test message",
			}, nil)
			req := create_request(message)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.Set(handlers.CurrentUserKey, alice)
			ctx.SetParamValues("1000")
			ctx.SetParamNames("chat")

			err := Create(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
