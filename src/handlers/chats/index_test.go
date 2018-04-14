package chats

import (
	"configuration"
	"db"
	"handlers"

	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

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
	if err := db.Create(&users); err != nil {
		t.Error(err)
	}

	if err := db.Create([]*db.Chat{
		&db.Chat{
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
		},
		&db.Chat{
			Name: "Chat one#001",
			Members: []*db.Member{
				&db.Member{
					UserID: alice.ID,
				},
				&db.Member{
					UserID: users[3].ID,
				},
				&db.Member{
					UserID: users[4].ID,
				},
				&db.Member{
					UserID: users[5].ID,
				},
			},
		},
	}); err != nil {
		t.Error(err)
	}

	Convey("Chat index tests", t, func() {
		Convey("Request to server with chats", func() {})
		Convey("Request with wrong limits", func() {})
	})
}
