package main

import "ueckoken/plarail2021-soft-builtin/internal"

func main() {
	// gRPC serves at :1111
	s := internal.NewGrpcServer(1111)
	s.StartServer()
}
