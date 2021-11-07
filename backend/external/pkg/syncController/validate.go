package syncController

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"ueckoken/plarail2021-soft-external/pkg/servo"
	"ueckoken/plarail2021-soft-external/pkg/stationNameId"
	"ueckoken/plarail2021-soft-external/spec"
)

type Validator struct {
	Stations []Station `yaml:"stations"`
}
type Station struct {
	Station EachStation `yaml:"station"`
}
type EachStation struct {
	Name   string   `yaml:"name"`
	Points []string `yaml:"points"`
	Rules  []Rule   `yaml:"rules"`
}
type Rule struct {
	On  []string `yaml:"on,omitempty"`
	Off []string `yaml:"off,omitempty"`
}

const (
	UNDEFINED = 0
	ALLOW     = 1
	DENY      = 2
)

//go:embed embed/stationRule.yml
var confFile []byte

func NewRouteValidator() *Validator {
	v := new(Validator)
	err := yaml.Unmarshal(confFile, v)
	if err != nil {
		panic(err)
	}
	return v
}
func (v *Validator) Validate(u StationState, ss []StationState) error {
	targetSta, err := v.getValidateTarget(u)
	if err != nil {
		return err
	}

	// バリデート対象外
	if reflect.DeepEqual(targetSta, &Station{}) {
		return nil
	}

	var ok [][]int
	// 置き替え後に正常
	id, err := searchIndex(u.StationID, ss)
	if err != nil {
		return err
	}
	ss[id] = StationState{servo.StationState{
		StationID: u.StationID,
		State:     u.State,
	}}
	for _, rule := range targetSta.Station.Rules {
		isOnOk := UNDEFINED
		isOffOk := UNDEFINED
		if rule.On == nil {
			isOnOk = ALLOW
		}
		for _, onRule := range rule.On {
			onId, err := stationNameId.Name2Id(onRule)
			if err != nil {
				return err
			}
			for _, kvsSta := range ss {
				if isOnOk == DENY {
					break
				}
				if kvsSta.StationID != onId {
					continue
				}
				// ルール合致
				if kvsSta.State == int32(spec.RequestSync_ON) {
					isOnOk = ALLOW
				} else {
					isOnOk = UNDEFINED
				}
			}
		}
		if rule.Off == nil {
			isOffOk = ALLOW
		}
		for _, offRule := range rule.Off {
			offId, err := stationNameId.Name2Id(offRule)
			if err != nil {
				return err
			}
			for _, kvsSta := range ss {
				if isOffOk == DENY {
					break
				}
				if kvsSta.StationID != offId {
					continue
				}
				if kvsSta.State == int32(spec.RequestSync_OFF) {
					isOffOk = ALLOW
				} else if kvsSta.State != int32(spec.RequestSync_OFF) {
					isOffOk = DENY
				}
			}
		}
		ok = append(ok, []int{isOnOk, isOffOk})
	}
	if !isOk(ok) {
		n, _ := stationNameId.Id2Name(u.StationID)
		return fmt.Errorf("validation `%s` error\n", n)
	}
	return nil
}

func isOk(b [][]int) bool {
	for _, eachRule := range b {
		on := eachRule[0]
		off := eachRule[1]
		if on == ALLOW && off == ALLOW {
			return true
		}
	}
	return false
}
func (v *Validator) getValidateTarget(u StationState) (*Station, error) {
	targetSta := new(Station)
	for _, s := range v.Stations {
		for _, pointName := range s.Station.Points {
			id, err := stationNameId.Name2Id(pointName)
			if err != nil {
				return nil, err
			}
			if u.StationID == id {
				targetSta = &s
				break
			}
		}
	}
	return targetSta, nil
}

func searchIndex(id int32, ss []StationState) (int, error) {
	for i, s := range ss {
		if s.StationID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("index error\n")
}
