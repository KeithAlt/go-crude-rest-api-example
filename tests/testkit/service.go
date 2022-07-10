package testkit

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"log"
	"net/http"
	"time"
)

var serviceStarted bool

// XXX this works, but we should probably set up our test services to deploy independently across our network
// our tests maybe passing right now, but with enough latency they won't do so less the above change is implemented

// CheckService checks the state of our service
func CheckService() {
	if !serviceStarted {
		go startService()
		serviceStarted = true
		time.Sleep(time.Second * 1)
	}
}

// startService starts our service
func startService() {
	log.Println("[TEST-util] Starting service...")
	config.Set()
	client := *repository.Initialize()
	defer api.Serve(&client)
}

// KillService shuts down our service
func KillService() {
	_, err := http.Get(config.Host + "/kill")
	if err != nil {
		log.Fatal("An error occurred when killing our service: ", err)
		return
	}
	log.Println("[TEST-util] Test service killed")
}

// CheckStatusCode checks the returned status code to ensure it's what we expect it to be
func CheckStatusCode(resCode int, passCodes []int) bool {
	for _, code := range passCodes {
		if resCode == code {
			return true
		}
	}
	return false
}
