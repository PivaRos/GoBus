package structs

import (
	"encoding/xml"
	"time"
)

type ResponseData struct {
	XMLName xml.Name `json:"-" xml:"Siri"`
	Siri    Siri     `json:"Siri" xml:"Siri"`
}

type Siri struct {
	ServiceDelivery ServiceDelivery `json:"ServiceDelivery" xml:"ServiceDelivery"`
}

type ServiceDelivery struct {
	ResponseTimestamp         time.Time                `json:"ResponseTimestamp" xml:"ResponseTimestamp"`
	ProducerRef               string                   `json:"ProducerRef" xml:"ProducerRef"`
	ResponseMessageIdentifier string                   `json:"ResponseMessageIdentifier" xml:"ResponseMessageIdentifier"`
	RequestMessageRef         string                   `json:"RequestMessageRef" xml:"RequestMessageRef"`
	Status                    string                   `json:"Status" xml:"Status"`
	StopMonitoringDelivery    []StopMonitoringDelivery `json:"StopMonitoringDelivery" xml:"StopMonitoringDelivery"`
}

type StopMonitoringDelivery struct {
	Version            string               `json:"-" xml:"version,attr"`
	ResponseTimestamp  time.Time            `json:"ResponseTimestamp" xml:"ResponseTimestamp"`
	Status             string               `json:"Status" xml:"Status"`
	MonitoredStopVisit []MonitoredStopVisit `json:"MonitoredStopVisit" xml:"MonitoredStopVisit"`
}

type MonitoredStopVisit struct {
	RecordedAtTime          time.Time               `json:"RecordedAtTime" xml:"RecordedAtTime"`
	ItemIdentifier          string                  `json:"ItemIdentifier" xml:"ItemIdentifier"`
	MonitoringRef           string                  `json:"MonitoringRef" xml:"MonitoringRef"`
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney" xml:"MonitoredVehicleJourney"`
}

type MonitoredVehicleJourney struct {
	LineRef                  string                  `json:"LineRef" xml:"LineRef"`
	DirectionRef             string                  `json:"DirectionRef" xml:"DirectionRef"`
	FramedVehicleJourneyRef  FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef" xml:"FramedVehicleJourneyRef"`
	PublishedLineName        string                  `json:"PublishedLineName" xml:"PublishedLineName"`
	OperatorRef              string                  `json:"OperatorRef" xml:"OperatorRef"`
	DestinationRef           string                  `json:"DestinationRef" xml:"DestinationRef"`
	OriginAimedDepartureTime time.Time               `json:"OriginAimedDepartureTime" xml:"OriginAimedDepartureTime"`
	ConfidenceLevel          string                  `json:"ConfidenceLevel" xml:"ConfidenceLevel"`
	VehicleLocation          VehicleLocation         `json:"VehicleLocation" xml:"VehicleLocation"`
	Bearing                  string                  `json:"Bearing" xml:"Bearing"`
	Velocity                 string                  `json:"Velocity" xml:"Velocity"`
	VehicleRef               string                  `json:"VehicleRef" xml:"VehicleRef"`
	MonitoredCall            MonitoredCall           `json:"MonitoredCall" xml:"MonitoredCall"`
}

type FramedVehicleJourneyRef struct {
	DataFrameRef           string `json:"DataFrameRef" xml:"DataFrameRef"`
	DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef" xml:"DatedVehicleJourneyRef"`
}

type VehicleLocation struct {
	Longitude string `json:"Longitude" xml:"Longitude"`
	Latitude  string `json:"Latitude" xml:"Latitude"`
}

type MonitoredCall struct {
	StopPointRef        string    `json:"StopPointRef" xml:"StopPointRef"`
	Order               string    `json:"Order" xml:"Order"`
	ExpectedArrivalTime time.Time `json:"ExpectedArrivalTime" xml:"ExpectedArrivalTime"`
	DistanceFromStop    string    `json:"DistanceFromStop" xml:"DistanceFromStop"`
}
