package syncController

import (
	"sync"
	"testing"
	"ueckoken/plarail2021-soft-external/pkg/servo"
)

func (skvs *stationKVS) contain(ss StationState) bool {
	for _, s := range skvs.stations {
		if s == ss {
			return true
		}
	}
	return false
}

func TestSyncController_update(t *testing.T) {
	station1 := StationState{servo.StationState{StationID: 1, State: 1}}
	station2 := StationState{servo.StationState{StationID: 2, State: 1}}
	kvs := stationKVS{
		stations: nil,
		mtx:      sync.Mutex{},
	}
	kvs.update(station1)
	if !kvs.contain(station1) {
		t.Errorf("append failed")
	}

	// new station append
	kvs.update(station2)
	if !kvs.contain(station2) {
		t.Errorf("station add failed")
	}
	if len(kvs.stations) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("stations before update are not keeping")
	}
	if !kvs.contain(station1) && !kvs.contain(station2) {
		t.Errorf("station2 is not append with `update` method")
	}

	// update exist station data
	station1 = StationState{servo.StationState{StationID: 1, State: 0}}
	kvs.update(station1)
	if len(kvs.stations) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("not update station data")
	}
}
func TestSyncController_get(t *testing.T) {
	station1 := StationState{servo.StationState{StationID: 1, State: 1}}
	station2 := StationState{servo.StationState{StationID: 2, State: 1}}
	skvs := stationKVS{
		stations: nil,
		mtx:      sync.Mutex{},
	}
	// member is not exist
	station, err := skvs.get(0)
	if station != (StationState{}) {
		t.Errorf("'station' is expect for empty but not empty")
	}

	if err == nil {
		t.Errorf("'err' is expect not nil")
	} else if err.Error() != "Not found" {
		t.Errorf("err.Error() expect 'Not found' but return %e", err)
	}

	skvs = stationKVS{
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

func Test_stationKVS_update(t *testing.T) {
	type fields struct {
		stations  []StationState
		validator *Validator
	}
	type args struct {
		u StationState
	}
	var stations []StationState
	for i := 1; i < 20; i++ {
		stations = append(stations,
			StationState{servo.StationState{StationID: int32(i), State: 1}})
	}
	v := Validator{Stations: []Station{{
		EachStation{
			Name:   "chofu_kudari",
			Points: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
			Rules: []Rule{{
				On:  nil,
				Off: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
			}, {
				On:  []string{"chofu_s1"},
				Off: nil,
			}, {
				On:  []string{"chofu_s2"},
				Off: nil,
			},
			},
		},
	}}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		/*
			1:  "chofu_b1",
			2:  "chofu_b2",
			3:  "chofu_b3",
			4:  "chofu_s1",
			5:  "chofu_s2",
			6:  "chofu_s3",
			7:  "chofu_s4",
		*/
		{
			name: "全てOFFの状態でOFFにすることはルールの1つ目に従う",
			fields: fields{
				stations:  stations,
				validator: &v,
			},
			args: args{StationState{servo.StationState{
				StationID: 1, // chofu_b1
				State:     2, // off
			}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skvs := &stationKVS{
				stations:  tt.fields.stations,
				validator: tt.fields.validator,
			}
			if err := skvs.update(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
