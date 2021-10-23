package internal

import (
	"errors"
	"sync"
)

type StationKV struct{
	StationID int32
	State int32
}

type stationKVS struct{
	stations []StationKV
	mtx sync.Mutex
}

func (skvs *stationKVS) update(u StationKV){
	skvs.mtx.Lock()
	for _,s := range skvs.stations {
		if s.StationID == u.StationID {
			s.State = u.State
			skvs.mtx.Unlock()
			return
		}
	}
	skvs.stations = append(skvs.stations, u)
	skvs.mtx.Unlock()
	return
}
func (skvs *stationKVS) get(stationID int32) (station StationKV, err error){
	skvs.mtx.Lock()
	for _,s := range skvs.stations{
		if s.StationID == station.StationID {
			skvs.mtx.Unlock()
			return s,nil
		}
	}
	skvs.mtx.Unlock()
	return StationKV{}, errors.New("Not found")
}
type SyncController struct{
	ClientHandler2syncController chan StationKV
	SyncController2clientHandler chan StationKV
}

func (s SyncController)StartSyncController(){
	var kvs stationKVS
	for c := range s.ClientHandler2syncController{
		kvs.update(c)
		s.SyncController2clientHandler <- c
	}
}
