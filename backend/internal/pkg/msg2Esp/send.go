package msg2Esp

import (
	"encoding/json"
	"time"
	"ueckoken/plarail2021-soft-internal/pkg/Send2Json"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
)

type sendData struct {
	State string `json:"state"`
	Pin   int    `json:"pin"`
}
type Send2Node interface {
	Send() error
}
type send2node struct {
	Station  *station2espIp.StationDetail
	TimeOut  time.Duration
	sendData *sendData
}

func NewSend2Node(sta *station2espIp.StationDetail, state string, timeOut time.Duration) Send2Node {
	return &send2node{
		Station: sta,
		TimeOut: timeOut,
		sendData: &sendData{
			State: state,
			Pin:   sta.Pin,
		},
	}
}
func (s *send2node) Send() error {
	b, err := json.Marshal(s.sendData)
	if err != nil {
		return err
	}

	err = Send2Json.SendJson(s.Station.Address, b, s.TimeOut)
	if err != nil {
		return err
	}

	return nil
}
