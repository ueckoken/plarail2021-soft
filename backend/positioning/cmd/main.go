package main

import (
	"fmt"
	"log"
	"os"
	"ueckoken/plarail2021-soft-positioning/internal"
	"ueckoken/plarail2021-soft-positioning/pkg/gormHandler"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbenv := os.Getenv("DB")
	fmt.Println("DB:", dbenv)
	d, err := gorm.Open(postgres.Open(dbenv), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to sql: %s", err)
		os.Exit(1)
	}
	db := gormHandler.SQLHandler{DB: d}
	db.AutoMigrate(&trainState.State{})
	r := internal.NewPositionReceiver(db)
	r.StartPositionReceiver()
}
