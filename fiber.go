package oauth2wall

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/laravel/passport"
)

func InitFiberRoute(app *fiber.App) {
	app.Get("/auth/passport/authorize", passport.Authorize)
	app.Get("/auth/passport/callback", passport.Token)
	app.Use(passport.Authorization)
}
