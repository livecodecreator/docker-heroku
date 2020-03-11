package raspi

import "github.com/gorilla/websocket"

var status Status

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var commandBroadcast = make(chan commandResponse)
var commandClients = make(map[*websocket.Conn]bool)

var chatBroadcast = make(chan chatMessage)
var chatClients = make(map[*websocket.Conn]bool)
