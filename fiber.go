package oauth2wall

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/common"
	"github.com/wangsying/oauth2wall/laravel/passport"
	"github.com/wangsying/oauth2wall/osin/simple"
)

func NewOAuth2(app *fiber.App) OAuth2 {
	return &oAuth2{
		app: app,
	}
}

type OAuth2 interface {
	InitLaravelPassportRoute() OAuth2
	InitOsinSimpleRoute() OAuth2
	Middleware()
}

type oAuth2 struct {
	app  *fiber.App
	mids []func(*fiber.Ctx) error
}

func (f *oAuth2) Middleware() {
	f.app.Use(func(ctx *fiber.Ctx) error {
		fmt.Printf("\noauth middleware route: %s\n", ctx.Route().Path)
		for _, m := range f.mids {
			if m(ctx) == nil {
				return ctx.Next()
			}
		}

		return ctx.JSON(common.NewResult("用户权限验证中间件验证用户信息失败").SetSuccess(false))
	})
}

func (f *oAuth2) InitLaravelPassportRoute() OAuth2 {
	f.app.Get("/auth/passport/authorize", passport.Authorize)
	f.app.Get("/auth/passport/callback", passport.Callback)

	f.mids = append(f.mids, passport.Middleware)
	return f
}

func (f *oAuth2) InitOsinSimpleRoute() OAuth2 {
	f.app.Get("/auth/osin/authorize", simple.Authorize)
	f.app.Get("/auth/osin/callback", simple.Callback)
	f.mids = append(f.mids, passport.Middleware)

	return f
}
