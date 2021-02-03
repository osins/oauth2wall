package passport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wangsying/oauth2wall/common"
)

type User struct {
	ID       int
	Name     string
	NickName string
	Email    string
	Token    string
}

func GetUser(token string) *common.Result {
	u := &User{}
	req, err := http.NewRequest("GET", config.UserURL, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err).SetData(body)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return common.NewResult(err.Error()).SetSuccess(false).SetError(err).SetData(body)
	}

	u.Token = token
	return common.NewResultSuccess(u)
}
