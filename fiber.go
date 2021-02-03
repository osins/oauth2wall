package oauth2wall

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/laravel/passport"
)

type FiberOAuth interface {
	Init(app *fiber.App)
}
type fiberOAuth struct {
}

func (f *fiberOAuth) Init(app *fiber.App) {
	app.Get("/auth/passport/authorize", passport.Authorize)
	app.Get("/auth/passport/callback", passport.Token)
	app.Use(passport.Authorization)
}
