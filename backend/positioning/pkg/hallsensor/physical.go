package hallsensor

import (
	_ "embed"
	"github.com/ghodss/yaml"
)

//go:embed embed/hallsensor.yaml
var HallSensorSetting []byte

type PhysicalSensors struct {
	sensors map[string]PhysicalSensor
}

func NewPhysicalSensors() PhysicalSensors {
	var y SensorYaml
	yaml.Unmarshal(HallSensorSetting, &y)
	p := make(map[string]PhysicalSensor)
	ps := PhysicalSensors{sensors: p}
	for _, h := range y.Halls {
		ps.sensors[h.Name] = PhysicalSensor{name: h.Name, predict: h.Predict, nexts: h.Nexts}
	}
	return ps
}

type PhysicalSensor struct {
	name    string
	predict bool
	nexts   []string
}

type SensorYaml struct {
	Halls []struct {
		Name    string   `yaml:"name"`
		Predict bool     `yaml:"predict"`
		Nexts   []string `yaml:"nexts"`
	} `yaml:"halls"`
}
