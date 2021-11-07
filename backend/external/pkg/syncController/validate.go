package syncController

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v2"
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

//go:embed embed/stationRule.yml
var yamlFile []byte

func NewRouteValidator() *Validator {
	v := new(Validator)
	err := yaml.Unmarshal(yamlFile, v)
	if err != nil {
		panic(err)
	}
	return v
}
func (v *Validator) Validate(u StationState, ss *[]StationState) error {
	targetSta, err := v.getValidateTarget(u)
	if err != nil {
		return err
	}
	if targetSta == nil {
		// バリデート対象外
		return nil
	}
	ok := [][]bool{{false, false}}
	for _, rule := range targetSta.Station.Rules {
		isOnOk := false
		isOffOk := false
		for _, onRule := range rule.On {
			onId, err := stationNameId.Name2Id(onRule)
			if err != nil {
				return err
			}
			for _, kvsSta := range *ss {
				if kvsSta.StationID != onId {
					continue
				}
				// ルール合致
				if kvsSta.State == int32(spec.RequestSync_ON) {
					isOnOk = true
				}
			}
		}
		for _, offRule := range rule.Off {
			offId, err := stationNameId.Name2Id(offRule)
			if err != nil {
				return err
			}
			for _, kvsSta := range *ss {
				if kvsSta.StationID != offId {
					continue
				}
				if kvsSta.State == int32(spec.RequestSync_OFF) {
					isOffOk = true
				}
			}
		}
		ok = append(ok, []bool{isOnOk, isOffOk})
	}
	if !isOk(ok) {
		return fmt.Errorf("validation `%d` error\n", u.StationID)
	}
	return nil
}

func isOk(b [][]bool) bool {
	for _, eachRule := range b {
		on := eachRule[0]
		off := eachRule[1]
		if on && off {
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
