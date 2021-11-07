package syncController

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/envStore"
	"ueckoken/plarail2021-soft-external/pkg/servo"
)

type StationState struct {
	servo.StationState
}

type stationKVS struct {
	stations  []StationState
	mtx       sync.Mutex
	validator *Validator
}

func newStationKvs() *stationKVS {
	v := NewRouteValidator()
	skvs := stationKVS{validator: v}
	return &skvs
}
func (skvs *stationKVS) update(u StationState) error {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	err := skvs.validator.Validate(u, skvs.stations)
	if err != nil {
		return err
	}
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return nil
		}
	}
	skvs.stations = append(skvs.stations, u)
	return nil
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
	Environment                  *envStore.Env
}

func (s *SyncController) StartSyncController() {
	kvs := newStationKvs()
	go s.periodicallySync(kvs)
	s.triggeredSync(s.Environment, kvs)
}

func (s *SyncController) triggeredSync(e *envStore.Env, kvs *stationKVS) {
	for c := range s.ClientHandler2syncController {
		kvs.update(c)
		c2i := servo.NewCommand2Internal(c.StationState, e)
		err := c2i.Send()
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@:", err)
		s.SyncController2clientHandler <- c
	}
}

func (s *SyncController) periodicallySync(kvs *stationKVS) {
	ch := time.Tick(2 * time.Second)
	for range ch {
		fmt.Println("lockig")
		kvs.mtx.Lock()
		fmt.Println("locked")
		for _, st := range kvs.retrieve() {
			s.SyncController2clientHandler <- st
		}
		kvs.mtx.Unlock()
	}
}
