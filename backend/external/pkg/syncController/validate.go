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

type IValidator interface {
	Validate(state StationState, ss []StationState) error
}
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

type RuleSuite struct {
	On  int
	Off int
}

const (
	UNDEFINED = 0
	ALLOW     = 1
	DENY      = 2
)

//go:embed embed/stationRule.yml
var confFile []byte

func NewRouteValidator() IValidator {
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
	var ok []RuleSuite
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
		isOnOk, err := matchRule(rule.On, ss, int32(spec.RequestSync_ON))
		if err != nil {
			return err
		}
		isOffOk, err = matchRule(rule.Off, ss, int32(spec.RequestSync_OFF))
		if err != nil {
			return err
		}
		ok = append(ok, RuleSuite{On: isOnOk, Off: isOffOk})
	}

	if !allRuleOk(ok) {
		n, _ := stationNameId.Id2Name(u.StationID)
		return fmt.Errorf("validation `%s` error\n", n)
	}
	return nil
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

func (r *RuleSuite) isOk() bool {
	return r.On == ALLOW && r.Off == ALLOW
}
func allRuleOk(rules []RuleSuite) bool {
	for _, r := range rules {
		if r.isOk() {
			return true
		}
	}
	return false
}
func searchIndex(id int32, ss []StationState) (int, error) {
	for i, s := range ss {
		if s.StationID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("index error\n")
}

// matchRule
func matchRule(rules []string, ss []StationState, state int32) (status int, err error) {
	isSuiteRule := UNDEFINED
	if rules == nil {
		isSuiteRule = ALLOW
	}
	for _, rule := range rules {
		id, err := stationNameId.Name2Id(rule)
		if err != nil {
			return -1, err
		}
		for _, kvsSta := range ss {
			if isSuiteRule == DENY {
				break
			}
			if kvsSta.StationID != id {
				continue
			}
			// ルール合致
			if kvsSta.State == state {
				isSuiteRule = ALLOW
			} else {
				isSuiteRule = DENY
			}
			break
		}
	}
	return isSuiteRule, nil
}
