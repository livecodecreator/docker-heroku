package slack

const (
	eventTypeURLVerification = "url_verification"
	eventTypeCallback        = "event_callback"
	eventCallbackTypeMessage = "message"
	contentType              = "Content-Type"
	applicationJSON          = "application/json; charset=utf-8"
	authorizationHeader      = "Authorization"
	bearerFormat             = "Bearer %s"
	chatPostMessageEndpoint  = "https://slack.com/api/chat.postMessage"
)
