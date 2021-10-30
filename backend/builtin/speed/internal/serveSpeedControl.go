package internal

import (
	"context"
	pb "ueckoken/plarail2021-soft-builtin/spec"
)

type ServeSpeedControl struct {
	pb.UnimplementedSpeedServer
}

func NewServeSpeedControl() *ServeSpeedControl {
	return &ServeSpeedControl{}
}
func (s *ServeSpeedControl) ControlSpeed(ctx context.Context, ss *pb.SendSpeed) (*pb.StatusCode, error) {
	return &pb.StatusCode{}, nil
}
