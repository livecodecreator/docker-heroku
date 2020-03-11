package raspi

import (
	"bytes"
	"io/ioutil"
	"log"
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

func commandMessangeBroadcastLoop() {

	for {
		msg := <-commandBroadcast
		for client := range commandClients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("%v\n", err)
				client.Close()
				delete(commandClients, client)
			}
		}
	}
}

func chatMessangeBroadcastLoop() {

	for {
		msg := <-chatBroadcast
		for client := range chatClients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("%v\n", err)
				client.Close()
				delete(chatClients, client)
			}
		}
	}
}
