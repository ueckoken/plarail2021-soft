package train2IP

import (
	_ "embed"
	"gopkg.in/yaml.v2"
	"log"
)

//go:embed embed/train2IP.yml
var confFile []byte

type Name2Id struct {
	Trains []struct {
		Name string `yaml:"name"`
		IP   string `yaml:"ip"`
	} `yaml:"trains"`
}

func GetTable() Name2Id {
	var n2i Name2Id
	err := yaml.Unmarshal(confFile, n2i)
	if err != nil {
		log.Fatalln("train2IP parse failed", err)
	}
	return n2i
}
func (t *Name2Id) changeForm2() {}

func (t *Name2Id) SearchIp(name string) (ip string) {
	for _, i := range t.Trains {
		if name == i.Name {
			return i.IP
		}
	}
	return ""
}
