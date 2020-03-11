package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/livecodecreator/docker-heroku/common"
	"github.com/livecodecreator/docker-heroku/raspi"
	"github.com/livecodecreator/docker-heroku/server"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	b, err := readBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	if challengeRequestIfNeeded(w, r, b) {
		return
	}

	if eventCallbackRequestIfNeeded(w, r, b) {
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func challengeRequestIfNeeded(w http.ResponseWriter, r *http.Request, b []byte) bool {

	var req challengeRequest
	err := json.Unmarshal(b, &req)
	if err != nil {
		log.Printf("%v\n", err)
		return false
	}

	if req.Type != eventTypeURLVerification {
		return false
	}

	res := challengeResponse{Challenge: req.Challenge}
	d, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return true
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(d))
	return true
}

func eventCallbackRequestIfNeeded(w http.ResponseWriter, r *http.Request, b []byte) bool {

	var req eventCallbackRequest
	err := json.Unmarshal(b, &req)
	if err != nil {
		log.Printf("%v\n", err)
		return false
	}

	if req.Type != eventTypeCallback {
		return false
	}

	if req.Event.Type != eventCallbackTypeMessage {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("event.type is not %v\n", eventCallbackTypeMessage)
		log.Printf("event.type: %v¥n", req.Event.Type)
		return true
	}

	log.Printf("event.text: %v\n", req.Event.Text)

	if strings.Contains(req.Event.Text, "hey") && strings.Contains(req.Event.Text, "pi") {
		postStatus(w, r)
	}

	w.WriteHeader(http.StatusOK)
	return true
}

func postStatus(w http.ResponseWriter, r *http.Request) {

	ss := server.GetStatus()
	rs := raspi.GetStatus()

	srw := statusResponseWrapper{HelloTime: time.Now().Format(time.RFC3339)}

	srw.ServerStatus = statusResponse{
		CPU:      fmt.Sprintf("%.2f", ss.CPU),
		Disk:     fmt.Sprintf("%.2f", ss.Disk),
		Memory:   fmt.Sprintf("%.2f", ss.Memory),
		BootTime: time.Unix(int64(ss.BootTime), 0).Format(time.RFC3339),
		PostTime: ss.PostTime.Format(time.RFC3339),
	}

	srw.RaspiStatus = statusResponse{
		CPU:      fmt.Sprintf("%.2f", rs.CPU),
		Disk:     fmt.Sprintf("%.2f", rs.Disk),
		Memory:   fmt.Sprintf("%.2f", rs.Memory),
		BootTime: time.Unix(int64(rs.BootTime), 0).Format(time.RFC3339),
		PostTime: rs.PostTime.Format(time.RFC3339),
	}

	d, err := json.MarshalIndent(srw, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	req := raspiStatusRequest{
		Channel: common.Env.SlackChannel,
		Text:    "```" + string(d) + "```",
	}

	b, err := json.Marshal(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	res, err := PostStatus(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v\n", err)
		return
	}

	log.Printf("slack api response body: %v\n", string(res))
}
