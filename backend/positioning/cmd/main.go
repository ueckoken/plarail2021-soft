package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"ueckoken/plarail2022-positioning/internal"
	"ueckoken/plarail2022-positioning/pkg/gormHandler"
	"ueckoken/plarail2022-positioning/pkg/trainState"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbenv := os.Getenv("DB")
	fmt.Println("DB:", dbenv)
	var d *gorm.DB
	var err error
	for {
		d, err = gorm.Open(postgres.Open(dbenv), &gorm.Config{})
		if err != nil {
			log.Printf("error when connecting to sql: %s", err)
			time.Sleep(1 * time.Second)
			log.Panicf("re-connecting to sql")
			continue
		}
		break
	}
	db := gormHandler.SQLHandler{Db: d}
	db.Db.AutoMigrate(&trainState.State{})
	r := internal.NewPositionReceiver(db, internal.NewApplicationStatus())
	r.StartPositionReceiver()
}
