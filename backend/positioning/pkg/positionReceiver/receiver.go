package positionReceiver

import (
	"encoding/json"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-positioning/pkg/hallsensor"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"
)

type PositionReceiveHandler struct {
	registerReceivedPosition chan trainState.State
}

func NewPositionReceiverHandler(registerReceivedPosition chan trainState.State) PositionReceiveHandler {
	return PositionReceiveHandler{
		registerReceivedPosition: registerReceivedPosition,
	}
}

type ReceivedPosition struct {
	MacAddress string `json:"mac_address"`
	Pin        int    `json:"pin"`
	TrainId    int    `json:"train_id"`
}

func (p PositionReceiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Use POST"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var buf []byte
	r.Body.Read(buf)
	var receivedPosition ReceivedPosition
	json.Unmarshal(buf, &receivedPosition)
	h := hallsensor.NewEsp32PinSetting()
	name, err := h.Search(receivedPosition.MacAddress, receivedPosition.Pin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found such address or pin"))
		return
	}
	//TODO: validate trainId
	dat := trainState.State{
		TrainId:          receivedPosition.TrainId,
		HallSensorName:   name,
		FetchedTimeStump: time.Now(),
	}
	p.registerReceivedPosition <- dat
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("accepted"))
}
