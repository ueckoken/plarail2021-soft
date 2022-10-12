package internal

import (
	"time"
	"ueckoken/plarail2021-soft-positioning/pkg/hallsensor"
)

type ApplicationStatus struct {
	TrainStatuses  []TrainStatus
	HallSensorSpec hallsensor.PhysicalSensors
}

func NewApplicationStatus() (ApplicationStatus, error) {
	pss, err := hallsensor.NewPhysicalSensors()
	if err != nil {
		return ApplicationStatus{}, err
	}
	return ApplicationStatus{HallSensorSpec: pss}, nil
}

type TrainStatus struct {
	TrainId               int
	FetchedHallSensorName string
	FetchedTimeStump      time.Time
}
