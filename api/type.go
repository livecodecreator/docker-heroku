package api

type statusResponse struct {
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
	RequestLength int    `json:"requestLength"`
}
