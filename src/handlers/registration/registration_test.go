package registration

import (
	"configuration"
	"db"
	"handlers"

	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func request(user, password string) *http.Request {
	form := url.Values{}
	form.Add("name", user)
	form.Add("password", password)

	r := httptest.NewRequest(http.MethodPost, "/registration", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return r
}

func TestRegisterUser(t *testing.T) {
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

	Convey("User registration tests", t, func() {

		Convey("Register user", func() {
			req := request("user", "password")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := RegisterUser(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusCreated)

			user := &db.User{
				Name: "user",
			}

			err = db.Get(user)
			So(err, ShouldBeNil)
			So(*user.Password, ShouldEqual, "password")

			Convey("Register new user with the same nickname", func() {
				req := request("user", "password")
				rec := httptest.NewRecorder()
				ctx = e.NewContext(req, rec)

				err := RegisterUser(ctx)
				So(err, ShouldBeNil)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("Register user: empty nickname", func() {
			req := request("", "password")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := RegisterUser(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("Register user: empty password", func() {
			req := request("user", "")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := RegisterUser(ctx)
			So(err, ShouldBeNil)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
