package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func readBody(r *http.Request) ([]byte, error) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return b, err
	}
	defer r.Body.Close()

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	return b, err
}
