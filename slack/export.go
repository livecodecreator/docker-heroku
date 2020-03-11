package slack

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/livecodecreator/docker-heroku/common"
)

// PostStatus is
func PostStatus(b []byte) (string, error) {

	req, err := http.NewRequest(http.MethodPost, chatPostMessageEndpoint, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	authorizationValue := fmt.Sprintf(bearerFormat, common.Env.SlackToken)
	req.Header.Set(authorizationHeader, authorizationValue)
	req.Header.Set(contentType, applicationJSON)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return string(dat), nil
}
