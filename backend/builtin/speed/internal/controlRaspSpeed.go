package internal

import (
	"fmt"
	"log"
	"net/http"
)

type RaspberrySpeed struct {
	speed int32
}

var endPoint = "http://localhost:8081"

//var endPoint = "http://192.168.100.8:8085"

func NewRaspberrySpeed(speed int32) *RaspberrySpeed {
	return &RaspberrySpeed{speed: speed}
}

func (r *RaspberrySpeed) changeSpeed() error {
	url := fmt.Sprintf("%s/?speed=%d", endPoint, r.speed)
	log.Printf("changeSpeed try to send to %s", url)
	res, err := http.Get(url)
	log.Println("changeSpeed: res,err", res, "\n", err)
	if res == nil {
		return err
	}
	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf("GET Err is `%w` ;HTTP status is `%s`", err, res.Status)
	}
	return err
}
