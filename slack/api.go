package slack

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/livecodecreator/docker-heroku/common"
	"github.com/urfave/negroni"
)

// SetupAPI is
func SetupAPI(r *mux.Router) {

	sr := mux.NewRouter().StrictSlash(true)
	sr.HandleFunc("/slack", defaultHandler)

	r.PathPrefix("/slack").Handler(negroni.New(
		negroni.HandlerFunc(authMiddleware),
		negroni.Wrap(sr),
	))
}

func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	b, err := readBody(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	var req authRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	if req.Token != common.Env.SlackVerificationToken {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid token request rejected")
		return
	}

	next(w, r)
}
