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
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			skvs.mtx.Unlock()
			return
		}
	}
	skvs.stations = append(skvs.stations, u)
	skvs.mtx.Unlock()
	return
}
func (skvs *stationKVS) get(stationID int32) (station StationState, err error) {
	skvs.mtx.Lock()
	for _, s := range skvs.stations {
		if s.StationID == stationID {
			skvs.mtx.Unlock()
			return s, nil
		}
	}
	skvs.mtx.Unlock()
	return StationState{}, errors.New("Not found")
}

type SyncController struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
}

func (s *SyncController) StartSyncController() {
	var kvs stationKVS
	go func() {
		ch := time.Tick(2 * time.Second)
		for _ = range ch {
			kvs.mtx.Lock()
			for _, st := range kvs.stations {
				s.SyncController2clientHandler <- st
			}
			kvs.mtx.Unlock()
		}
	}()
	go func(){
		for c := range s.ClientHandler2syncController {
			kvs.update(c)
			s.SyncController2clientHandler <- c
		}
	}()
}
