package service

type authRequest struct {
	Token string `json:"token"`
}

type searchRequest struct {
	Keyword string `json:"keyword"`
}

type searchResponse struct {
	Keyword string `json:"keyword"`
}
