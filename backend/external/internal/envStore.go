package internal

import (
	"github.com/go-playground/validator"
	"github.com/vrischmann/envconfig"
	"time"
)

type hostnamePort string

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

type Env struct {
	ClientSideServer struct {
		Addr hostnamePort `default:"127.0.0.1:54321"`
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
