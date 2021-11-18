package addressChecker

import (
	_ "embed"
	"log"
	"net"

	"github.com/ghodss/yaml"
)

//go:embed embed/allowlist.yaml
var AllowListYaml []byte

type AddressChecker struct {
	allowlist []string
}

func (a AddressChecker) CheckIfOk(addr string) bool {
	for _, d := range a.allowlist {
		_, ipnet, err := net.ParseCIDR(d)
		if err != nil {
			log.Fatal()
		}
		ip, _, err := net.ParseCIDR(addr)
		if err != nil {
			log.Printf("failed to parse cidr: %s\n", addr)
		}
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

type allowList struct {
	Network []string `yaml:"network"`
}

func NewAddressChecker() (*AddressChecker, error) {
	t := allowList{}
	err := yaml.Unmarshal(AllowListYaml, &t)
	if err != nil {
		return nil, err
	}
	ret := AddressChecker{}
	for _, n := range t.Network {
		_, _, err := net.ParseCIDR(n)
		if err != nil {
			log.Fatalln(err)
		}
		ret.allowlist = append(ret.allowlist, n)
	}
	return &ret, nil
}
