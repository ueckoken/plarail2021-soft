package main

import (
	"ueckoken/plarail2021-soft-external/internal"
)

func main() {
	go internal.StartServer()
	go internal.StartSyncController()
	for {
	}
}
