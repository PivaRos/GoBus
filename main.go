package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	gobus "github.com/pivaros/GoBus/src"
)

func main() {
	log.Println("starting")
	serviceUri := "<your uri>"
	Key := "<your key>"
	client := http.Client{
		Transport: &http.Transport{},
		Timeout:   10 * time.Second,
	}
	gobus, err := gobus.InitGoBus(serviceUri, Key, client)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("init success")
	result, monitorErr := gobus.MonitoringRef(1)
	if monitorErr != nil {
		log.Panicln(monitorErr)
	}
	for index, value := range result.ServiceDelivery.StopMonitoringDelivery {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
