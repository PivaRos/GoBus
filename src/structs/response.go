package structs

import (
	"encoding/xml"
	"time"
)

type ResponseData struct {
	XMLName         xml.Name `json:"-" xml:"Siri"`
	ServiceDelivery struct {
		ResponseTimestamp         time.Time `json:"ResponseTimestamp" xml:"ResponseTimestamp"`
		ProducerRef               string    `json:"ProducerRef" xml:"ProducerRef"`
		ResponseMessageIdentifier string    `json:"ResponseMessageIdentifier" xml:"ResponseMessageIdentifier"`
		RequestMessageRef         string    `json:"RequestMessageRef" xml:"RequestMessageRef"`
		Status                    string    `json:"Status" xml:"Status"`
		StopMonitoringDelivery    []struct {
			Version            string    `json:"-" xml:"version,attr"`
			ResponseTimestamp  time.Time `json:"ResponseTimestamp" xml:"ResponseTimestamp"`
			Status             string    `json:"Status" xml:"Status"`
			MonitoredStopVisit []struct {
				RecordedAtTime          time.Time `json:"RecordedAtTime" xml:"RecordedAtTime"`
				ItemIdentifier          string    `json:"ItemIdentifier" xml:"ItemIdentifier"`
				MonitoringRef           string    `json:"MonitoringRef" xml:"MonitoringRef"`
				MonitoredVehicleJourney struct {
					LineRef                 string `json:"LineRef" xml:"LineRef"`
					DirectionRef            string `json:"DirectionRef" xml:"DirectionRef"`
					FramedVehicleJourneyRef struct {
						DataFrameRef           string `json:"DataFrameRef" xml:"DataFrameRef"`
						DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef" xml:"DatedVehicleJourneyRef"`
					} `json:"FramedVehicleJourneyRef" xml:"FramedVehicleJourneyRef"`
					PublishedLineName        string    `json:"PublishedLineName" xml:"PublishedLineName"`
					OperatorRef              string    `json:"OperatorRef" xml:"OperatorRef"`
					DestinationRef           string    `json:"DestinationRef" xml:"DestinationRef"`
					OriginAimedDepartureTime time.Time `json:"OriginAimedDepartureTime" xml:"OriginAimedDepartureTime"`
					ConfidenceLevel          string    `json:"ConfidenceLevel" xml:"ConfidenceLevel"`
					VehicleLocation          struct {
						Longitude string `json:"Longitude" xml:"Longitude"`
						Latitude  string `json:"Latitude" xml:"Latitude"`
					} `json:"VehicleLocation" xml:"VehicleLocation"`
					Bearing       string `json:"Bearing" xml:"Bearing"`
					Velocity      string `json:"Velocity" xml:"Velocity"`
					VehicleRef    string `json:"VehicleRef" xml:"VehicleRef"`
					MonitoredCall struct {
						StopPointRef        string    `json:"StopPointRef" xml:"StopPointRef"`
						Order               string    `json:"Order" xml:"Order"`
						ExpectedArrivalTime time.Time `json:"ExpectedArrivalTime" xml:"ExpectedArrivalTime"`
						DistanceFromStop    string    `json:"DistanceFromStop" xml:"DistanceFromStop"`
					} `json:"MonitoredCall" xml:"MonitoredCall"`
				} `json:"MonitoredStopVisit" xml:"MonitoredStopVisit"`
			} `json:"StopMonitoringDelivery" xml:"StopMonitoringDelivery"`
		} `json:"ServiceDelivery" xml:"ServiceDelivery"`
	} `json:"Siri" xml:"-"`
}
