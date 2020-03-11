package slack

type authRequest struct {
	Token string `json:"token"`
}

type statusRequest struct {
	CPU      string `json:"cpu"`
	Disk     string `json:"disk"`
	Memory   string `json:"memory"`
	BootTime string `json:"bootTime"`
	PostTime string `json:"postTime"`
}

type statusResponse struct {
	CPU      string `json:"cpu"`
	Disk     string `json:"disk"`
	Memory   string `json:"memory"`
	BootTime string `json:"bootTime"`
	PostTime string `json:"postTime"`
}

type statusResponseWrapper struct {
	HelloTime    string         `json:"helloTime"`
	ServerStatus statusResponse `json:"serverStatus"`
	RaspiStatus  statusResponse `json:"raspiStatus"`
}

type challengeRequest struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type challengeResponse struct {
	Challenge string `json:"challenge"`
}

type eventCallbackEvent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type eventCallbackRequest struct {
	Token string             `json:"token"`
	Type  string             `json:"type"`
	Event eventCallbackEvent `json:"event"`
}

type raspiStatusRequest struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
