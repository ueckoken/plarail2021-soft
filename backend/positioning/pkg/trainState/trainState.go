package trainState

import (
	"time"
)

type States struct {
	States []State
}

type State struct {
	TrainId          int
	HallSensorName   string
	FetchedTimeStump time.Time
}

type TrainState interface {
	Store(State) error
	FetchLatest(trainId string) error
}
