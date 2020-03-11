package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

func incrementRequestCount() {

	atomic.AddUint32(&requestCount, 1)
}

func readBody(r *http.Request) ([]byte, error) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return b, err
	}
	defer r.Body.Close()

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	return b, err
}
