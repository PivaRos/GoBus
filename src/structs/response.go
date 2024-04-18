package structs

import (
	"encoding/xml"
	"time"
)

type ResponseData struct {
	XMLName xml.Name `json:"-"`
	Siri    Siri     `json:"Siri"`
}

type Siri struct {
	ServiceDelivery ServiceDelivery `json:"ServiceDelivery"`
}

type ServiceDelivery struct {
	ResponseTimestamp         time.Time                `json:"ResponseTimestamp"`
	ProducerRef               string                   `json:"ProducerRef"`
	ResponseMessageIdentifier string                   `json:"ResponseMessageIdentifier"`
	RequestMessageRef         string                   `json:"RequestMessageRef"`
	Status                    string                   `json:"Status"`
	StopMonitoringDelivery    []StopMonitoringDelivery `json:"StopMonitoringDelivery"`
}

type StopMonitoringDelivery struct {
	Version            string                `json:"-"attr"`
	ResponseTimestamp  time.Time             `json:"ResponseTimestamp"`
	Status             string                `json:"Status"`
	MonitoredStopVisit *[]MonitoredStopVisit `json:"MonitoredStopVisit,omitempty"`
	ErrorCondition     *ErrorCondition       `json:"ErrorCondition,omitempty"`
}

type MonitoredStopVisit struct {
	RecordedAtTime          time.Time               `json:"RecordedAtTime"`
	ItemIdentifier          string                  `json:"ItemIdentifier"`
	MonitoringRef           string                  `json:"MonitoringRef"`
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney"`
}

type MonitoredVehicleJourney struct {
	LineRef                  string                  `json:"LineRef"`
	DirectionRef             string                  `json:"DirectionRef"`
	FramedVehicleJourneyRef  FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef"`
	PublishedLineName        string                  `json:"PublishedLineName"`
	OperatorRef              string                  `json:"OperatorRef"`
	DestinationRef           string                  `json:"DestinationRef"`
	OriginAimedDepartureTime time.Time               `json:"OriginAimedDepartureTime"`
	ConfidenceLevel          string                  `json:"ConfidenceLevel"`
	VehicleLocation          VehicleLocation         `json:"VehicleLocation"`
	Bearing                  string                  `json:"Bearing"`
	Velocity                 string                  `json:"Velocity"`
	VehicleRef               string                  `json:"VehicleRef"`
	MonitoredCall            MonitoredCall           `json:"MonitoredCall"`
}

type FramedVehicleJourneyRef struct {
	DataFrameRef           string `json:"DataFrameRef"`
	DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef"`
}

type VehicleLocation struct {
	Longitude string `json:"Longitude"`
	Latitude  string `json:"Latitude"`
}

type MonitoredCall struct {
	StopPointRef        string    `json:"StopPointRef"`
	Order               string    `json:"Order"`
	ExpectedArrivalTime time.Time `json:"ExpectedArrivalTime"`
	DistanceFromStop    string    `json:"DistanceFromStop"`
}

type ErrorCondition struct {
	OtherError  *OtherError `json:"OtherError,omitempty"omitempty"`
	Description string      `json:"Description"`
}

type OtherError struct {
	ErrorText string `json:"ErrorText"`
}
