package internal

import (
	"context"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type ControlServer struct {
	pb.UnimplementedControlServer
}

func (c *ControlServer) Command2Internal(ctx context.Context, req *pb.RequestSync) (*pb.ResponseSync, error) {
	return &pb.ResponseSync{Response: pb.ResponseSync_UNKNOWN}, nil
}
