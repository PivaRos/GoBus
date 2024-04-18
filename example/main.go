package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	gobus "github.com/pivaros/GoBus/src"
)

func main() {
	log.Println("starting")
	serviceUri := "<Your Uri>"
	Key := "<Your Key>"
	client := http.Client{
		Transport: &http.Transport{},
		Timeout:   10 * time.Second,
	}
	rdb_Options := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	options := gobus.GoBusOptions{
		ServiceUri:  serviceUri,
		ServiceKey:  Key,
		Client:      client,
		Rdb_Options: rdb_Options,
		StaleTime:   50 * time.Second,
	}
	gobus, err := gobus.InitGoBus(options)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("init success")
	result, monitorErr := gobus.MonitoringRef("1")
	if monitorErr != nil {
		log.Panicln(monitorErr)
	}
	for index, value := range *result.Siri.ServiceDelivery.StopMonitoringDelivery[0].MonitoredStopVisit {
		jsonData, err := json.MarshalIndent(value.MonitoredVehicleJourney, "", "    ")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf("Index: %d, Value: %d\n", index, string(jsonData))
	}

}
