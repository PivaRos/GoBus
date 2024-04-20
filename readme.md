# GoBus

GoBus is a Go SDK designed to interface with the Siri-Standard to enhance and simplify bus tracking. This application provides users with real-time updates on bus locations, helping improve the efficiency and reliability of public transportation.

The SIRI (Service Interface for Real Time Information) standard is a protocol designed to facilitate the exchange of real-time information regarding public transportation services between different information systems.

## Features

- **Real-Time Bus Tracking**:Track bus locations in real-time using the Siri International system.
- **Easy To Configure**: Configuration via environment variables for ease of deployment.
- **Built-in Caching**: Built-in Redis support for caching and managing data state.
- **Super Fast Integration**: Easy setup and integration with Go applications.
- **Program-Friendly Interface**: Simple and intuitive SDK to easily track buses.
- **Notifications**: Get notified about bus arrival times and delays.

## Getting Started v1.0.1

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:

- [Redis](https://redis.io/) server (for caching responses and managing state).
- [Go](https://golang.org/dl/) (version 1.16 or higher)

### Installation

Install GoBus using go get:

```go
go get github.com/pivaros/GoBus
```

### Configuration

GoBus uses environment variables for configuration. Set the following variables in your .env file:

- SIRI_SERVICE_URI: The URI of the SIRI service endpoint.
- SIRI_SERVICE_KEY: The authentication key for accessing the SIRI service.

Example of .env file:

```plaintext
SIRI_SERVICE_URI=http://example.com/siri
SIRI_SERVICE_KEY=your_service_key
```

### Usage

##### Initializing the SDK

First, import the necessary packages and initialize the GoBus client:

```go
package main

import (
    "log"
    "net/http"
    "time"
    "github.com/go-redis/redis/v8"
    "github.com/pivaros/GoBus"
)

func main() {

    options := gobus.GoBusOptions{
        Client:      http.Client{
        Transport: &http.Transport{},
        Timeout: <Your Timeout>,   // Customize the timeout as needed
        },
        Rdb_Options: redis.Options{
        Addr:     <Your Address>,  // Redis server address
        Password: <Your Password>, // Redis password if set
        DB:       <Your DB>,       // Redis database number
        },
        StaleTime:<Your Duration>, // Duration after which the data is considered stale
    }

    gobus, err := gobus.InitGoBus(options)
    if err != nil {
        log.Panicln(err)
    }

    // Now you can use the gobus instance to interact with the SIRI services
}
```

### Fetching Real-Time Data

To fetch real-time data, use the MonitoringRef function:

```go
result, monitorErr := gobus.MonitoringRef("<your_monitoring_ref>")
if monitorErr != nil {
    //Handle Error
    log.Panicln(monitorErr)
}

// Handle the result
// Note: The result is automatically decoded into a Go struct
log.Println("log part of the resonse", result.ServiceDelivery.Status)
```

### Example Of Usage

```go
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
	client := http.Client{
	.../etc
	}
	rdb_Options := redis.Options{
	.../etc
	}
	options := gobus.GoBusOptions{
	.../options
	}

	gobus, err := gobus.InitGoBus(options)
	if err != nil {
		log.Panicln(err)
	}
	result, monitorErr := gobus.MonitoringRef("1")
	if monitorErr != nil {
		log.Panicln(monitorErr)
	}
	//print every entry
	for index, value := range *result.Siri.ServiceDelivery.StopMonitoringDelivery[0].MonitoredStopVisit {
		jsonData, err := json.MarshalIndent(value.MonitoredVehicleJourney, "", "    ")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf("Index: %d, Value: %d\n", index, string(jsonData))
	}
}
```

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

Fork the Project
Create your Feature Branch (git checkout -b feature/AmazingFeature)
Commit your Changes (git commit -m 'Add some AmazingFeature')
Push to the Branch (git push origin feature/AmazingFeature)
Open a Pull Request

## License

Distributed under the MIT License. See LICENSE for more information.

## Contact

Daniel Gurbin - danielgavgurbin@gmail.com

Project Link: https://github.com/PivaRos/GoBus
