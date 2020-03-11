package server

import "time"

// Status is
type Status struct {
	CPU      float64   `json:"cpu"`
	Disk     float64   `json:"disk"`
	Memory   float64   `json:"memory"`
	BootTime uint64    `json:"bootTime"`
	PostTime time.Time `json:"postTime"`
}
