package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/livecodecreator/docker-heroku/common"
	"github.com/livecodecreator/docker-heroku/raspi"
	"github.com/urfave/negroni"
)

// SetupAPI is
func SetupAPI(r *mux.Router) {

	sr := mux.NewRouter().StrictSlash(true)
	sr.HandleFunc("/service/search", searchHandler)

	r.PathPrefix("/service").Handler(negroni.New(
		negroni.HandlerFunc(authMiddleware),
		negroni.Wrap(sr),
	))
}

func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var err error
	var b []byte

	if b, err = readBody(r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	var req authRequest
	if err = json.Unmarshal(b, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	if req.Token != common.Env.ServiceToken {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid token request rejected")
		return
	}

	next(w, r)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var b []byte

	if b, err = readBody(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	var req searchRequest
	if err = json.Unmarshal(b, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	res := searchResponse{Keyword: req.Keyword}
	if b, err = json.Marshal(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	raspi.CommandPush("takeshot")

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
