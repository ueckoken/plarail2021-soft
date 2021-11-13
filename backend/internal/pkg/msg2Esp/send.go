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
	Angle int    `json:"angle"`
}
type send2nodeExistAngle struct {
	Station  *station2espIp.StationDetail
	TimeOut  time.Duration
	sendData *sendData
}

type sendDataNoAngle struct {
	State string `json:"state"`
	Pin   int    `json:"pin"`
}

type send2nodeNoAngle struct {
	send2nodeExistAngle
	sendData *sendDataNoAngle
	Station  *station2espIp.StationDetail
	TimeOut  time.Duration
}

type Send2Node interface {
	Send() error
}

func NewSend2Node(sta *station2espIp.StationDetail, state string, angle int, timeOut time.Duration) Send2Node {
	if sta.IsAngleDefined() {
		return &send2nodeExistAngle{
			Station: sta,
			TimeOut: timeOut,
			sendData: &sendData{
				State: state,
				Pin:   sta.Pin,
				Angle: angle,
			},
		}
	} else {
		return &send2nodeNoAngle{
			Station: sta,
			TimeOut: timeOut,
			sendData: &sendDataNoAngle{
				State: state,
				Pin:   sta.Pin,
			},
		}
	}
}

func (s *send2nodeExistAngle) Send() error {
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
