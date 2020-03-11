package common

import (
	"os"
)

// Env is
var Env = env{
	Port:                   os.Getenv("PORT"),
	RaspiToken:             os.Getenv("RASPI_TOKEN"),
	ServiceToken:           os.Getenv("SERVICE_TOKEN"),
	MackerelAPIKey:         os.Getenv("MACKEREL_API_KEY"),
	SlackToken:             os.Getenv("SLACK_TOKEN"),
	SlackChannel:           os.Getenv("SLACK_CHANNEL"),
	SlackVerificationToken: os.Getenv("SLACK_VERIFICATION_TOKEN"),
}
