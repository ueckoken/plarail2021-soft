package gormHandler

import (
	"gorm.io/gorm"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"
)

type SQLHandler struct {
	*gorm.DB
}

func (s SQLHandler) Store(state trainState.State) {
	s.Create(state)
}

func (s SQLHandler) FetchFromTrainId(trainId int) (state trainState.States) {
	s.Where("train_id = ?", trainId).Find(&state.States)
	return state
}

func (s SQLHandler) FetchFromHallSensorName(hallId string) (state trainState.States){
	s.Where("hall_sensor_name = ?", hallId).Find(&state.States)
	return state
}