package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type ControlServer struct {
	pb.UnimplementedControlServer
	env      *Env
	Stations *station2espIp.Stations
}

func (c *ControlServer) Command2Internal(ctx context.Context, req *pb.RequestSync) (*pb.ResponseSync, error) {
	sta, err := c.unpackStations(req.GetStation())
	if err != nil {
		return nil, err
	}
	s2n := NewSend2Node(sta, c.unpackState(req.State), c.env)
	err = s2n.Send2Esp()

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "sender err %s; not connected to Node", err.Error())
	}
	return &pb.ResponseSync{Response: pb.ResponseSync_SUCCESS}, nil
}

func (c *ControlServer) unpackStations(req *pb.Stations) (*station2espIp.StationDetail, error) {
	s, ok := pb.Stations_StationId_name[int32(req.GetStationId())]
	if !ok {
		return nil, fmt.Errorf("station: %s do not define in proto file\n", req.String())
	}
	sta, err := c.Stations.SearchStation(s)
	if err != nil {
		return nil, fmt.Errorf("station %s do not define in yaml file\n", s)
	}
	return sta, nil
}
func (c *ControlServer) unpackState(state pb.RequestSync_State) string {
	return state.String()
}
