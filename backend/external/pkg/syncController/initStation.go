package syncController

import (
	_ "embed"
	"gopkg.in/yaml.v2"
	"log"
)

type InitRule struct {
	Stations []StationInitRule `yaml:"stations"`
}
type StationInitRule struct {
	Name  string `yaml:"name"`
	State int    `yaml:"state"`
}

//go:embed embed/stationInit.yml
var initRuleFile []byte

func NewInitializeRule() *InitRule {
	r := new(InitRule)
	err := yaml.Unmarshal(initRuleFile, r)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}
