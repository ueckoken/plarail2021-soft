package internal

import (
	"net/http"
	"ueckoken/plarail2021-soft-positioning/pkg/gormHandler"
	"ueckoken/plarail2021-soft-positioning/pkg/positionReceiver"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"
)

type PositionReceiver struct {
	db gormHandler.SQLHandler
}
func NewPositionReceiver(db gormHandler.SQLHandler) PositionReceiver{
	return PositionReceiver{db}
}

func (pos PositionReceiver) StartPositionReceiver() {
	c := make(chan trainState.State)
	p := positionReceiver.NewPositionReceiverHandler(c, pos.db)
	http.Handle("/position", p)
	http.ListenAndServe(":8080", nil)
}
