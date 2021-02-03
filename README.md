# oauth2wall 关于OAuth2的基于fiber的web框架下的登录验证库

该库的开发初衷是为了要对接Laravel的Passport登录验证功能，故，目前仅支持基于web框架fiber的Laravel的Passport验证。

几个相关oauth的路由如下：
```
/auth/passport/authorize
/auth/passport/callback
```

关键入口：
https://github.com/wangsying/oauth2wall/blob/5172bc88d897bb89554c6ad44998e82b2af6fe8e/fiber.go

初始化代码如下：

```

  import (
    "github.com/gofiber/fiber/v2"
    "github.com/wangsying/oauth2wall"
  )

  app := fiber.New()

  fiberInit(app)

  oauth2wall.NewFiberOAuth2().Init(app)

  app.Listen(":8087")

  return app
 ```
