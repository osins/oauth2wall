// Package passport implements the OAuth2 protocol for authenticating users
// through passport.
package passport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/common"
	"golang.org/x/oauth2"
)

const SESSION_STATE_KEY = "oauthStateString"

func GenOAuthState(ctx *fiber.Ctx) (string, *common.Result) {
	randomString := common.Hex(16)
	session, err := sessionStore.Get(ctx)
	defer session.Save()
	if err != nil {
		return "", common.NewResult("session 回话加载失败.").SetSuccess(false)
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

func Authorize(ctx *fiber.Ctx) error {
	s, r := GenOAuthState(ctx)
	if r != nil {
		return r.Error
	}

	return ctx.Redirect(oAuth2Config.AuthCodeURL(s), http.StatusTemporaryRedirect)
}

func Token(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if s, r := GetOAuthState(ctx); r != nil && state != s {
		ctx.JSON(common.NewResult(fmt.Sprintf("state 验证失败")).SetSuccess(false))
		return r.Error
	}

	sessionStore.Storage.Delete(SESSION_STATE_KEY)

	token, err := oAuth2Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		ctx.JSON(common.NewResult("获取用户信息失败").SetSuccess(false))
		return err
	}

	ctx.JSON(GetUser(token.AccessToken))
	return nil
}

func Authorization(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	result := GetUser(token)
	if !result.Success {
		return result.Error
	}

	return ctx.JSON(result)
}

func GetUser(token string) *common.Result {
	u := &User{}
	req, err := http.NewRequest("GET", config.UserURL, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err).SetData(u)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err).SetData(u)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err).SetData(u)
	}

	u.Token = token
	return common.NewResultSuccess(u)
}
