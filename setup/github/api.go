package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type API struct {
	host     string
	username string
	token    string
}

func NewApi(host, username, token string) API {
	return API{host, username, token}
}

func (a API) get(path string, model any) error {
	uri, err := url.JoinPath("https://", a.host, path)
	if err != nil {
		return fmt.Errorf("erroring creating url (%v)", err)
	}
	log.Debugf("GET: %v %v", uri, a.username)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return fmt.Errorf("erroring creating request (%v)", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	if a.token != "" {
		return fmt.Errorf("NOT SETUP")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making api request (%v)", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("bad api response (%v)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body (%v)", err)
	}

	err = json.Unmarshal(body, model)
	if err != nil {
		return fmt.Errorf("error reading body (%v)", err)
	}

	return nil
}
