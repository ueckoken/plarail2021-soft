package grpcMock

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
	pb "ueckoken/plarail2022-internal/spec"
)

type GrpcMock struct {
	pb.UnimplementedControlServer
}

func (m *GrpcMock) Command2Internal(ctx context.Context, req *pb.RequestSync) (*pb.ResponseSync, error) {
	rand.Seed(time.Now().UnixMicro())
	if rand.Float64() > 0.8 {
		var c codes.Code
		r := rand.Float64()
		if r < 0.3 {
			c = codes.OutOfRange
		} else if r < 0.6 {
			c = codes.NotFound
		} else {
			c = codes.Unimplemented
		}
		return nil, status.Errorf(c, "this msg comes MOCK\n")
	}
	return &pb.ResponseSync{Response: pb.ResponseSync_SUCCESS}, nil
}
