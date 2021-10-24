package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	pb "ueckoken/plarail2021-soft-external/spec"
)

const (
	address    = "127.0.0.1:12345"
	timeoutSec = 1
	UNKNOWN    = "UNKNOWN"
	SUCCESS    = "SUCCESS"
	FAILED     = "FAILED"
)

type Command2Internal struct {
	station *StationState
}

func NewCommand2Internal(state StationState) *Command2Internal {
	return &Command2Internal{station: &state}
}

func (c2i *Command2Internal) sendRaw() (*pb.ResponseSync, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewControlClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSec*time.Second)
	defer cancel()
	r, err := c.Command2Internal(ctx, c2i.convert2pb())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c2i *Command2Internal) trapErr(rs *pb.ResponseSync, grpcErr error) error {
	if rs == nil { // gRPC error occur
		return fmt.Errorf("gRPC Err: %w", grpcErr)
	} else { // check Response Status
		switch rs.Response.String() {
		case UNKNOWN:
			return fmt.Errorf("%w; gRPC Response Error. Status is %s", grpcErr, UNKNOWN)
		case SUCCESS:
			return nil
		case FAILED:
			return fmt.Errorf("%w; gRPC Response Error. Status is %s", grpcErr, FAILED)
		default:
			return fmt.Errorf("%w; Unknown error is occured", grpcErr)
		}
	}
}

func (c2i *Command2Internal) Send() error {
	return c2i.trapErr(c2i.sendRaw())
}

func (c2i *Command2Internal) convert2pb() *pb.RequestSync {
	return &pb.RequestSync{
		Station: &pb.Stations{StationId: pb.Stations_StationId(c2i.station.StationID)},
		State:   pb.RequestSync_State(c2i.station.State),
	}
}
