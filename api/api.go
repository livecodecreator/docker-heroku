package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/livecodecreator/docker-heroku/common"
	"github.com/livecodecreator/docker-heroku/raspi"
	"github.com/livecodecreator/docker-heroku/service"
	"github.com/livecodecreator/docker-heroku/slack"
	"github.com/urfave/negroni"
)

// ListenAndServe is
func ListenAndServe() {

	r := mux.NewRouter().StrictSlash(true)

	r.Path("/").HandlerFunc(defaultHandler)
	r.Path("/status").Methods(http.MethodGet).HandlerFunc(defaultHandler)

	slack.SetupAPI(r)
	raspi.SetupAPI(r)
	service.SetupAPI(r)

	l := negroni.NewLogger()
	l.SetFormat("{{.Status}} {{.Duration}} {{.Request}}")
	n := negroni.New(negroni.NewRecovery(), l, negroni.HandlerFunc(requestCountMiddleware))

	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%s", common.Env.Port))
}

func requestCountMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	incrementRequestCount()
	next(w, r)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	b, err := readBody(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	sr := statusResponse{
		Status:        "OK",
		RequestLength: len(b),
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	d, err := json.Marshal(sr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}
