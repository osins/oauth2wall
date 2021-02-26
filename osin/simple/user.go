package simple

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/osins/oauth2wall/common"
)

type User struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
	ExpiresIn    int    `json:"ExpiresIn"`
	Scope        string `json:"Scope"`
	Raw          struct {
		ID        string    `json:"Id"`
		ClientID  string    `json:"ClientId"`
		Username  string    `json:"Username"`
		EMail     string    `json:"EMail"`
		Mobile    string    `json:"Mobile"`
		CreatedAt time.Time `json:"CreatedAt"`
	} `json:"UserData"`
}

func GetUser(token string) *common.Result {
	req, err := http.NewRequest("GET", config.UserURL, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return common.NewResult(fmt.Sprintf("connect oauth2 server error: %s", err.Error())).SetSuccess(false).SetError(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return common.NewResult(fmt.Sprintf("read return auth body error: %s", err.Error())).SetSuccess(false).SetError(err).SetData(body)
	}

	var b interface{}
	if err := json.Unmarshal(body, &b); err != nil {
		return common.NewResult(fmt.Sprintf("json unmarshal body error: %s", err.Error())).SetSuccess(false).SetError(err).SetData(body)
	}

	return common.NewResult("get user info success").SetSuccess(true).SetData(b)
}
