package raspi

import "time"

type authRequest struct {
	Token string `json:"token"`
}

type statusRequest struct {
	CPU      float64 `json:"cpu"`
	Disk     float64 `json:"disk"`
	Memory   float64 `json:"memory"`
	BootTime uint64  `json:"bootTime"`
}

type commandRequest struct {
	Message string `json:"message"`
}

type commandResponse struct {
	Command string `json:"command"`
}

type chatMessage struct {
	Message string `json:"message"`
}

// Status is
type Status struct {
	CPU      float64   `json:"cpu"`
	Disk     float64   `json:"disk"`
	Memory   float64   `json:"memory"`
	BootTime uint64    `json:"bootTime"`
	PostTime time.Time `json:"postTime"`
}
