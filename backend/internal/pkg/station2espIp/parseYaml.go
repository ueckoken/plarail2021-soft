package station2espIp

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Stations struct {
	Stations []Station `yaml:"stations"`
}
type Station struct {
	Station StationDetail `yaml:"station"`
}
type StationDetail struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Pin     int    `yaml:"pin"`
}

//go:embed embed/station2espIp.yml
var ConfigFile []byte

func NewStations() (*Stations, error) {
	t := Stations{}
	err := yaml.Unmarshal(ConfigFile, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Stations) SearchStation(name string) (*StationDetail, error) {
	for _, sta := range s.Stations {
		if name == sta.Station.Name {
			return &sta.Station, nil
		}
	}
	return nil, fmt.Errorf("station %s not found", name)
}
