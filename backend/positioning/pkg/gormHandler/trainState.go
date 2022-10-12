package gormHandler

import (
	"gorm.io/gorm"
	"ueckoken/plarail2022-positioning/pkg/trainState"
)

type SQLHandler struct {
	Db *gorm.DB
}

func (s SQLHandler) Store(state trainState.State) {
	s.Db.Create(state)
}

func (s SQLHandler) FetchFromTrainId(trainId int) (state trainState.States) {
	s.Db.Where("train_id = ?", trainId).Order("fetched_time_stump asc").Find(&state.States)
	return state
}

func (s SQLHandler) FetchFromHallSensorName(hallId string) (state trainState.States) {
	s.Db.Where("hall_sensor_name = ?", hallId).Order("fetched_time_stump asc").Find(&state.States)
	return state
}

func (s SQLHandler) Fetch(trainId int, hallId string) (state trainState.States) {
	s.Db.Where("hall_sensor_name = ?", hallId).Where("train_id = ?", trainId).Order("fetched_time_stump asc").Find(&state.States)
	return state
}
