package internal

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "ueckoken/plarail2021-soft-builtin/spec"
)

type ServeSpeedControl struct {
	pb.UnimplementedSpeedServer
}

func NewServeSpeedControl() *ServeSpeedControl {
	return &ServeSpeedControl{}
}
func (s *ServeSpeedControl) ControlSpeed(ctx context.Context, ss *pb.SendSpeed) (*pb.StatusCode, error) {
	rs := NewRaspberrySpeed(ss.Speed)
	err := rs.changeSpeed()
	if err != nil {
		return nil,
			status.Errorf(codes.Unavailable, "grpc; `%e`", err)
	}
	return &pb.StatusCode{Code: pb.StatusCode_SUCCESS}, nil
}
