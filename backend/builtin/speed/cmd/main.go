package cmd

import "ueckoken/plarail2021-soft-builtin/internal"

func main() {
	s := internal.NewGrpcServer(1111)
	s.StartServer()
}
