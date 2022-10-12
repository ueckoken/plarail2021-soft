package positionReceiver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2022-positioning/pkg/addressChecker"
	"ueckoken/plarail2022-positioning/pkg/hallsensor"
	"ueckoken/plarail2022-positioning/pkg/trainState"
)

type PositionReceiveHandler struct {
	registerReceivedPosition chan trainState.State
	checker                  *addressChecker.AddressChecker
}

func NewPositionReceiverHandler(registerReceivedPosition chan trainState.State) PositionReceiveHandler {
	checker, err := addressChecker.NewAddressChecker()
	if err != nil {
		log.Fatalln(err)
	}
	return PositionReceiveHandler{
		registerReceivedPosition: registerReceivedPosition,
		checker:                  checker,
	}
}

type ReceivedPosition struct {
	MacAddress string `json:"mac_address"`
	Pin        int    `json:"pin"`
	TrainId    int    `json:"train_id"`
}

func (p PositionReceiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := getClientIp(r)
	if !p.checker.CheckIfOk(ip) {
		w.WriteHeader(http.StatusForbidden)
	}
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

func getClientIp(r *http.Request) string {
	f := r.Header.Get("X-FORWARDED-FOR")
	if f != "" {
		return f
	}
	return r.RemoteAddr
}
