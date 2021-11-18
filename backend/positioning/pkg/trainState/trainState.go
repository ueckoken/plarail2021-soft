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

type PositionAndSpeed struct {
	State
	Speed float64
}
