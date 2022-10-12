package internal

import (
	"time"
	"ueckoken/plarail2022-positioning/pkg/hallsensor"
)

type ApplicationStatus struct {
	TrainStatuses  []TrainStatus
	HallSensorSpec hallsensor.PhysicalSensors
}

func NewApplicationStatus() ApplicationStatus {
	return ApplicationStatus{HallSensorSpec: hallsensor.NewPhysicalSensors()}
}

type TrainStatus struct {
	TrainId               int
	FetchedHallSensorName string
	FetchedTimeStump      time.Time
}
