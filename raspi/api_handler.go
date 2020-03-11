package raspi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {

	b, err := readBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v\n", err)
		return
	}

	var req statusRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	status = Status{
		CPU:      req.CPU,
		Disk:     req.Disk,
		Memory:   req.Memory,
		BootTime: req.BootTime,
		PostTime: time.Now(),
	}

	w.WriteHeader(http.StatusOK)
}

func getCommandHandler(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	commandClients[ws] = true

	for {
		var req commandRequest
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Printf("%v\n", err)
			delete(commandClients, ws)
			break
		}
	}
}

func chatMessageHandler(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	chatClients[ws] = true

	for {
		var msg chatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("%v\n", err)
			delete(chatClients, ws)
			break
		}
		chatBroadcast <- msg
	}
}
