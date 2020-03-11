package scheduler

import (
	"github.com/livecodecreator/docker-heroku/mackerel"
	"github.com/livecodecreator/docker-heroku/raspi"
	"github.com/livecodecreator/docker-heroku/server"
	"github.com/robfig/cron"
)

// StartScheduler is
func StartScheduler() {

	c := cron.New()
	c.AddFunc("@every 1m", server.UpdateStatus)
	c.AddFunc("@every 1m", mackerel.PostStatus)
	c.AddFunc("@every 5m", mackerel.PostRequestCount)
	c.AddFunc("@every 30s", raspi.ChatMessageKeepalive)
	c.AddFunc("@every 30s", raspi.CommandMessageKeepalive)
	c.Start()
}
