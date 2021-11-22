package internal

import (
	"context"
	"log"
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
	log.Printf("gRPC received...speed: `%v`, train: `%v`", ss.Speed, ss.Train)
	_ = rs.changeSpeed()
	//if err != nil {
	//	return nil,
	//		status.Errorf(codes.Unavailable, "grpc; `%e`", err)
	//}
	return &pb.StatusCode{Code: pb.StatusCode_SUCCESS}, nil
}
