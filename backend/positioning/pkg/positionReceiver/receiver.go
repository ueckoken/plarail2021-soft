package positionReceiver

import (
	"fmt"
	"net/http"
	"strconv"
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
	ip_forwarded := r.Header.Get("X-FORWARDED-FOR")
	if ip_forwarded != "" {
		ip = ip_forwarded
	} else {
		ip = r.RemoteAddr
	}
	h := hallsensor.NewEsp32PinSetting()
	pin_i, err := strconv.Atoi(pin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid pin"))
	}
	name, err := h.Search(ip, pin_i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not found such address or pin"))
	}
	trainId := r.URL.Query().Get("trainId")
	//TODO: validate trainId
	fmt.Println(name,trainId)
}
