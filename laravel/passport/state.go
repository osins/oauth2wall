package passport

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/common"
)

const SESSION_STATE_KEY = "oauthStateString"

func GenOAuthState(ctx *fiber.Ctx) (string, *common.Result) {
	randomString := common.Hex(16)
	session, err := sessionStore.Get(ctx)
	defer session.Save()
	if err != nil {
		return "", common.NewResult("session 回话加载失败.").SetSuccess(false).SetError(err)
	}

	session.Set(SESSION_STATE_KEY, randomString)
	return common.HmacSha256(randomString, oAuth2Config.ClientID+oAuth2Config.ClientSecret+oAuth2Config.Endpoint.AuthURL), nil
}

func GetOAuthState(ctx *fiber.Ctx) (string, *common.Result) {
	session, err := sessionStore.Get(ctx)
	defer session.Save()
	if err != nil {
		return "", common.NewResult("session 回话加载失败.").SetSuccess(false).SetError(err)
	}

	randomString := fmt.Sprintf("%v", session.Get(SESSION_STATE_KEY))
	oauthStateString := common.HmacSha256(randomString, oAuth2Config.ClientID+oAuth2Config.ClientSecret+oAuth2Config.Endpoint.AuthURL)
	return oauthStateString, nil
}
