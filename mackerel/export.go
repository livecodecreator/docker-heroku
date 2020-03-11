package mackerel

import (
	"encoding/json"
	"log"
	"time"

	"github.com/livecodecreator/docker-heroku/api"
	"github.com/livecodecreator/docker-heroku/raspi"
	"github.com/livecodecreator/docker-heroku/server"
)

// PutMetricDataPoint is
func PutMetricDataPoint(name string, value float64) error {

	now := time.Now().Unix()

	mdp := MetricDataPoints{
		MetricDataPoint{
			Name:  name,
			Time:  now,
			Value: value,
		},
	}

	return PutMetricDataPoints(mdp)
}

// PutMetricDataPoints is
func PutMetricDataPoints(mdp MetricDataPoints) error {

	b, err := json.Marshal(mdp)
	if err != nil {
		return err
	}

	dat, err := postRequest(b)
	if err != nil {
		return err
	}

	log.Printf("mackerel api response body: %v\n", dat)

	return nil
}

// PostStatus is
func PostStatus() {

	now := time.Now().Unix()

	serverStatus := server.GetStatus()
	raspiStatus := raspi.GetStatus()

	mdp := MetricDataPoints{
		MetricDataPoint{
			Name:  "Server.CPU",
			Time:  now,
			Value: serverStatus.CPU,
		},
		MetricDataPoint{
			Name:  "Server.Memory",
			Time:  now,
			Value: serverStatus.Memory,
		},
		MetricDataPoint{
			Name:  "Raspi.CPU",
			Time:  now,
			Value: raspiStatus.CPU,
		},
		MetricDataPoint{
			Name:  "Raspi.Memory",
			Time:  now,
			Value: raspiStatus.Memory,
		},
	}

	err := PutMetricDataPoints(mdp)
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// PostRequestCount is
func PostRequestCount() {

	rc := api.SwapRequestCount()
	err := PutMetricDataPoint("Service.Request", rc)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
