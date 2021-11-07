package positionReceiver

import (
	"net/http"
	"strconv"
	"time"
	"ueckoken/plarail2021-soft-positioning/pkg/gormHandler"
	"ueckoken/plarail2021-soft-positioning/pkg/hallsensor"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"
)

type PositionReceiveHandler struct {
	registerReceivedPosition chan trainState.State
	db                       gormHandler.SQLHandler
}

func NewPositionReceiverHandler(registerReceivedPosition chan trainState.State, db gormHandler.SQLHandler) PositionReceiveHandler {
	return PositionReceiveHandler{
		registerReceivedPosition: registerReceivedPosition,
		db:                       db,
	}
}

func (p PositionReceiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pin := r.URL.Query().Get("pin")
	var ip string
	ipForwarded := r.Header.Get("X-FORWARDED-FOR")
	if ipForwarded != "" {
		ip = ipForwarded
	} else {
		ip = r.RemoteAddr
	}
	h := hallsensor.NewEsp32PinSetting()
	pinI, err := strconv.Atoi(pin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid pin"))
	}
	name, err := h.Search(ip, pinI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found such address or pin"))
	}
	trainId := r.URL.Query().Get("trainId")
	trainIdI, err := strconv.Atoi(trainId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid trainId"))
	}
	//TODO: validate trainId
	dat := trainState.State{
		TrainId:          trainIdI,
		HallSensorName:   name,
		FetchedTimeStump: time.Now(),
	}
	p.db.Store(dat)
	p.registerReceivedPosition <- dat
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("accepted"))
}
