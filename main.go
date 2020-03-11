package main

import (
	"log"

	"github.com/livecodecreator/docker-heroku/api"
	"github.com/livecodecreator/docker-heroku/scheduler"
)

func main() {

	log.SetFlags(log.Lshortfile)
	scheduler.StartScheduler()
	api.ListenAndServe()
}
