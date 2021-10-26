package internal

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/vrischmann/envconfig"
	"strconv"
	"time"
)

type hostnamePort string
type Port int32

func (t *hostnamePort) Unmarshal(s string) error {
	addr := struct {
		Addr string `validate:"hostname_port"`
	}{Addr: s}
	err := validator.New().Struct(&addr)
	if err != nil {
		return err
	}
	*t = hostnamePort(s)
	return nil
}
func (t *hostnamePort) String() string {
	return string(*t)
}
func (p *Port) Unmarshal(s string) error {
	d, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if !(0 < d && d <= 65535) {
		return fmt.Errorf("Port range failed; Port: %d\n", d)
	}
	*p = Port(d)
	return nil
}

type Env struct {
	ClientSideServer struct {
		Port Port `default:"54321"`
	}
	Grpc struct {
		Addr       hostnamePort
		TimeoutSec time.Duration `default:"1s"`
		//SslCertPath string
	}
}

func GetEnv() *Env {
	var env Env
	if err := envconfig.Init(&env); err != nil {
		panic(err)
	}
	return &env
}
