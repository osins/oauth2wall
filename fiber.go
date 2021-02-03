package oauth2wall

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/laravel/passport"
)

func NewFiberOAuth2() FiberOAuth2 {
	return &fiberOAuth2{}
}

type FiberOAuth2 interface {
	Init(app *fiber.App)
}
type fiberOAuth2 struct {
}

func (f *fiberOAuth2) Init(app *fiber.App) {
	app.Get("/auth/passport/authorize", passport.Authorize)
	app.Get("/auth/passport/callback", passport.Token)
	app.Use(passport.Authorization)
}
