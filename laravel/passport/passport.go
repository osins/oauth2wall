// Package passport implements the OAuth2 protocol for authenticating users
// through passport.
package passport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wangsying/oauth2wall/common"
	"golang.org/x/oauth2"
)

func Authorize(ctx *fiber.Ctx) error {
	return ctx.Redirect(oAuth2Config.AuthCodeURL(oauthStateString), http.StatusTemporaryRedirect)
}

func Token(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if state != oauthStateString {
		return ctx.SendString("invalid oauth state")
	}

	token, err := oAuth2Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return ctx.SendString("code exchange failed: " + err.Error())
	}

	result := GetUser(token.AccessToken)

	return ctx.JSON(result)
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
