package raspi

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

	sr.HandleFunc("/raspi/command", getCommandHandler).Methods(http.MethodGet)
	sr.HandleFunc("/raspi/status", statusHandler).Methods(http.MethodPost)

	r.PathPrefix("/raspi/command").Handler(negroni.New(
		negroni.HandlerFunc(authHeaderMiddleware),
		negroni.Wrap(sr),
	))

	r.PathPrefix("/raspi/status").Handler(negroni.New(
		negroni.HandlerFunc(authBodyMiddleware),
		negroni.Wrap(sr),
	))

	go commandMessangeBroadcastLoop()

	// http://localhost:8080/public/raspi/chat/
	fs := http.FileServer(http.Dir("/root/go/src/github.com/livecodecreator/docker-heroku/public/raspi/chat"))
	sp := http.StripPrefix("/public/raspi/chat/", fs)
	r.HandleFunc("/public/raspi/message", chatMessageHandler)
	r.PathPrefix("/public/raspi/chat").Handler(sp)
	go chatMessangeBroadcastLoop()
}

func authHeaderMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.Header.Get(headerAuthorization) != common.Env.RaspiToken {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid token request rejected")
		return
	}

	next(w, r)
}

func authBodyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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

	if req.Token != common.Env.RaspiToken {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid token request rejected")
		return
	}

	next(w, r)
}
