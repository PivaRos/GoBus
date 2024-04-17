package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	gobus, err := gobus.InitGoBus(serviceUri, Key, client)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("init success")
	lineRef := "10523"
	result, monitorErr := gobus.MonitoringRef("1", &lineRef)
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
