# oauth2wall 关于OAuth2的基于fiber的web框架下的登录验证库

该库的开发初衷是为了要对接Laravel的Passport登录验证功能，故，目前仅支持基于web框架fiber的Laravel的Passport验证。

几个相关oauth的路由如下：
```
/auth/passport/authorize
/auth/passport/callback
```

关键入口：<br>
https://github.com/wangsying/oauth2wall/blob/5172bc88d897bb89554c6ad44998e82b2af6fe8e/fiber.go

初始化代码如下：

```

  import (
    "github.com/gofiber/fiber/v2"
    "github.com/wangsying/oauth2wall"
  )

  app := fiber.New()

	fiberInit(app)
	fmt.Printf("new oauth.")
	oauth2wall.NewOAuth2(app).InitLaravelPassportRoute().InitOsinSimpleRoute().Middleware()

	app.Get("/api/user", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.Locals("user"))
	})

  app.Listen(":8087")

  return app
 ```
 
 .env 相关配置:
 
 ```
LARAVEL_PASSPORT_CLIENT_ID=92a23f69-3c5e-459a-999c-f97761ffec0a
LARAVEL_PASSPORT_CLIENT_SECRET=Eg9SNWpMER6gRVmHFbRfab0EEbK9Y9HsaGY4Hx2i
LARAVEL_PASSPORT_ENDPOINT=http://localhost:8080
LARAVEL_PASSPORT_REDIRECT_URL=http://localhost:8087/auth/passport/callback

OSIN_CLIENT_ID=1234
OSIN_CLIENT_SECRET=aabbccdd
OSIN_ENDPOINT=http://localhost:14000
OSIN_REDIRECT_URL=http://localhost:8087/auth/osin/callback
```
