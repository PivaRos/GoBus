package gobus

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pivaros/GoBus/src/structs"
	"github.com/pivaros/GoBus/src/utils"
)

var Gobus *GoBus

type GoBus struct {
	client   *http.Client
	instance utils.Instance
}

func InitGoBus(ServiceUri string, ServiceKey string, client http.Client) (*GoBus, error) {

	instance, err := utils.CreateInstance(ServiceUri, ServiceKey)
	if err != nil {
		return Gobus, err
	}

	Gobus = &GoBus{
		client:   &client,
		instance: *instance,
	}
	return Gobus, nil
}

func (bus *GoBus) MonitoringRef(id interface{}) (*structs.ResponseData, error) {
	log.Println("here0")
	request := bus.instance
	params := request.URL.Query()
	switch id := id.(type) {
	case int:
		params.Add("MonitoringRef", strconv.Itoa(id))
	case string:
		params.Add("MonitoringRef", id)
	}
	request.URL.RawQuery = params.Encode()
	log.Println(request.Request.URL)
	r, httpErr := bus.client.Do(request.Request)
	log.Println("here0.5")
	if httpErr != nil {
		return nil, httpErr
	}
	log.Println("here1")
	var data structs.ResponseData
	decoder := json.NewDecoder(r.Body)
	parseErr := decoder.Decode(&data)
	if parseErr != nil {
		return nil, parseErr
	}
	return &data, nil
}
