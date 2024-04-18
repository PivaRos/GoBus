package gobus

import (
	"encoding/json"
	"errors"
	"net/http"

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

func (bus *GoBus) MonitoringRef(MonitoringId string, LineId *string) (*structs.ResponseData, error) {
	request := bus.instance
	params := request.URL.Query()
	params.Add("MonitoringRef", MonitoringId)

	if LineId != nil {
		params.Add("LineRef", *LineId)
	}
	request.URL.RawQuery = params.Encode()

	r, httpErr := bus.client.Do(request.Request)
	if httpErr != nil {
		return nil, httpErr
	}
	var data structs.ResponseData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	if data.Siri.ServiceDelivery.Status == "false" {
		return nil, errors.New(data.Siri.ServiceDelivery.ErrorCondition.Description)
	}
	if data.Siri.ServiceDelivery.StopMonitoringDelivery[0].Status == "false" {
		return nil, errors.New(data.Siri.ServiceDelivery.StopMonitoringDelivery[0].ErrorCondition.Description)
	}
	return &data, nil
}
