package raspi

// GetStatus is
func GetStatus() Status {

	return status
}

// CommandPush is
func CommandPush(command string) {

	res := commandResponse{Command: command}
	commandBroadcast <- res
}

// CommandMessageKeepalive is
func CommandMessageKeepalive() {

	res := commandResponse{Command: commandKeepalive}
	commandBroadcast <- res
}

// ChatMessageKeepalive is
func ChatMessageKeepalive() {

	msg := chatMessage{Message: "keepalive"}
	chatBroadcast <- msg
}
