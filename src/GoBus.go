package gobus

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/pivaros/GoBus/src/structs"
	"github.com/pivaros/GoBus/src/utils"
)

var Gobus *GoBus

type GoBus struct {
	client   *http.Client
	instance utils.Instance
	rdb      *redis.Client
	options  *GoBusOptions
}

type GoBusOptions struct {
	Client      http.Client
	Rdb_Options redis.Options
	StaleTime   time.Duration
}

func InitGoBus(options GoBusOptions) (*GoBus, error) {
	godotenv.Load()
	serviceUri := os.Getenv("SIRI_SERVICE_URI")
	serviceKey := os.Getenv("SIRI_SERVICE_KEY")
	if serviceUri == "" {
		return nil, errors.New("SIRI_SERVICE_URI was not found in .env file")
	}
	if serviceKey == "" {
		return nil, errors.New("SIRI_SERVICE_KEY was not found in .env file")
	}

	instance, err := utils.CreateInstance(serviceUri, serviceKey)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&options.Rdb_Options)
	rError := utils.CheckRedisConnection(context.TODO(), rdb)
	if rError != nil {
		return nil, rError
	}
	Gobus = &GoBus{
		client:   &options.Client,
		instance: *instance,
		rdb:      rdb,
		options:  &options,
	}

	return Gobus, nil
}

func (bus *GoBus) MonitoringRef(MonitoringId string) (*structs.ResponseData, error) {
	result, redisErr := bus.rdb.Get(context.TODO(), MonitoringId).Result()
	if redisErr == redis.Nil {
		request := bus.instance.Clone(context.Background())
		params := request.URL.Query()
		params.Add("MonitoringRef", MonitoringId)
		request.URL.RawQuery = params.Encode()

		r, httpErr := bus.client.Do(request)
		if httpErr != nil {
			return nil, httpErr
		}
		Data := structs.ResponseData{}
		err := json.NewDecoder(r.Body).Decode(&Data)
		if err != nil {
			return nil, err
		}
		byteArray, marshalErr := json.Marshal(Data)
		if marshalErr != nil {
			return nil, marshalErr
		}
		bus.rdb.Set(context.TODO(), MonitoringId, byteArray, bus.options.StaleTime)
		return handleResponse(&Data)
	} else if redisErr != nil {
		return nil, redisErr
	} else {
		//does have entry in redis
		Data := structs.ResponseData{}
		marshalErr := json.Unmarshal([]byte(result), &Data)
		if marshalErr != nil {
			return nil, marshalErr
		}
		return handleResponse(&Data)
	}
}

func handleResponse(data *structs.ResponseData) (*structs.ResponseData, error) {
	if data.Siri.ServiceDelivery.Status == "false" {
		return nil, errors.New(data.Siri.ServiceDelivery.ErrorCondition.Description)
	}
	if data.Siri.ServiceDelivery.StopMonitoringDelivery[0].Status == "false" {
		return nil, errors.New(data.Siri.ServiceDelivery.StopMonitoringDelivery[0].ErrorCondition.Description)
	}
	//if response does not consist of errors return the response struct
	return data, nil
}
