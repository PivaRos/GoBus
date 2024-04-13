package gobus

import (
	"net/http"
	"strconv"

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

func (bus *GoBus) MonitoringRef(id interface{}) {
	url := bus.instance.URL
	params := url.Query()
	switch id := id.(type) {
	case int:
		params.Add("MonitoringRef", strconv.Itoa(id))
	case string:
		params.Add("MonitoringRef", id)

	}
	url.RawQuery = params.Encode()
	
		
}
