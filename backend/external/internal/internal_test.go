package internal

import (
	"errors"
	"sync"
	"testing"
	pb "ueckoken/plarail2021-soft-external/spec"
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
	station1 := StationState{StationID: 1, State: 1}
	station2 := StationState{StationID: 2, State: 1}
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
	station1 = StationState{StationID: 1, State: 0}
	kvs.update(station1)
	if len(kvs.stations) != 2 {
		t.Errorf("append failed")
	}
	if !kvs.contain(station1) {
		t.Errorf("not update station data")
	}
}

func TestSyncController_get(t *testing.T) {
	station1 := StationState{StationID: 1, State: 1}
	station2 := StationState{StationID: 2, State: 1}
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

func TestCommand2Internal_trapResponseGrpcErr(t *testing.T) {
	// status 0 UNKNOWN; 1 SUCCESS; 2 FAILED
	testPatterns := []struct {
		grpcErr        error
		inResponseSync *pb.ResponseSync
		expectErrMsg   string
	}{
		{
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(0)},
			expectErrMsg:   "gRPC Err: TEST ERROR; gRPC Response status is UNKNOWN",
		}, {
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(1)},
			expectErrMsg:   "gRPC Err: TEST ERROR; gRPC Response status is SUCCESS",
		}, {
			grpcErr:        errors.New("TEST ERROR"),
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(2)},
			expectErrMsg:   "gRPC Err: TEST ERROR; gRPC Response status is FAILED",
		}, {
			grpcErr:        nil,
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(0)},
			expectErrMsg:   "gRPC Err: %!w(<nil>); gRPC Response status is UNKNOWN",
		}, {
			grpcErr:        nil,
			inResponseSync: &pb.ResponseSync{Response: pb.ResponseSync_Response(2)},
			expectErrMsg:   "gRPC Err: %!w(<nil>); gRPC Response status is FAILED",
		},
	}

	for _, tp := range testPatterns {
		t.Run(tp.expectErrMsg,
			func(t *testing.T) {
				err := trapResponseGrpcErr(tp.inResponseSync, tp.grpcErr)
				if err == nil {
					t.Errorf("Expect err occured, but not occured.")
				} else if err.Error() != tp.expectErrMsg {
					t.Errorf("err format failed. Actual err is: %e", err)
				}
			},
		)
	}

	// Normal
	var grpcErr error = nil
	rs := &pb.ResponseSync{Response: pb.ResponseSync_Response(1)}
	err := trapResponseGrpcErr(rs, grpcErr)
	if err != nil {
		t.Errorf("Expect err is NOT occured, but occured. : %e", err)
	}
}
