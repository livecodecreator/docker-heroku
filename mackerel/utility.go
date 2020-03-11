package mackerel

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/livecodecreator/docker-heroku/common"
)

func postRequest(b []byte) (string, error) {

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Set(contentType, applicationJSON)
	req.Header.Set(xAPIKey, common.Env.MackerelAPIKey)

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
