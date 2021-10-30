package internal

import (
	"errors"
	"sync"
	"time"
)

type StationState struct {
	StationID int32
	State     int32
}

type stationKVS struct {
	stations []StationState
	mtx      sync.Mutex
}

func (skvs *stationKVS) update(u StationState) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return
		}
	}
	skvs.stations = append(skvs.stations, u)
	return
}
func (skvs *stationKVS) get(stationID int32) (station StationState, err error) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for _, s := range skvs.stations {
		if s.StationID == stationID {
			return s, nil
		}
	}
	return StationState{}, errors.New("Not found")
}

func (skvs *stationKVS) retrieve() []StationState {
	return skvs.stations
}

type SyncController struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
	Environment                  *Env
}

func (s *SyncController) StartSyncController() {
	var kvs stationKVS
	go func() {
		ch := time.Tick(2 * time.Second)
		for _ = range ch {
			kvs.mtx.Lock()
			for _, st := range kvs.retrieve() {
				s.SyncController2clientHandler <- st
			}
			kvs.mtx.Unlock()
		}
	}()
	go func() {
		for c := range s.ClientHandler2syncController {
			kvs.update(c)
			s.SyncController2clientHandler <- c
		}
	}()
}
