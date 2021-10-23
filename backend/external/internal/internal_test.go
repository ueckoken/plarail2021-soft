package internal

import (
	"sync"
	"testing"
)

func TestSyncController_get(t *testing.T) {
	station1 := StationState{StationID: 1, State: 1}
	station2 := StationState{StationID: 2, State: 1}
	skvsNil := stationKVS{
		stations: nil,
		mtx:      sync.Mutex{},
	}
	// member is not exist
	station, err := skvsNil.get(0)
	if station != (StationState{}) {
		t.Errorf("'station' is expect for empty but not empty")
	}

	if err == nil {
		t.Errorf("'err' is expect not nil")
	} else if err.Error() != "Not found" {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}

	skvs := stationKVS{
		stations: []StationState{station1, station2},
		mtx:      sync.Mutex{},
	}
	station, err = skvs.get(1)
	if station != station1 {
		t.Errorf("'station1' is expect but called station%d", station.StationID)
	}
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}

	station, err = skvs.get(2)
	if station != station2 {
		t.Errorf("'station2' is expect but called station%d", station.StationID)
	}
	if err != nil {
		t.Errorf("return err is not nil: %e", err)
	}

	// test for call 'get' not exist record
	station, err = skvs.get(3)
	if station != (StationState{}) {
		t.Errorf("'station' is expect for empty but called station%d", station.StationID)
	}
	if err == nil {
		t.Errorf("expect err but return nil")
	} else if err.Error() != "Not found" {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}
}
