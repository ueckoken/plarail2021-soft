package internal

import (
	"context"
	"fmt"
	"ueckoken/plarail2021-soft-internal/pkg"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type ControlServer struct {
	pb.UnimplementedControlServer
	env      *Env
	Stations *pkg.Stations
}

func (c *ControlServer) Command2Internal(ctx context.Context, req *pb.RequestSync) (*pb.ResponseSync, error) {
	sta, err := c.unpackStations(req.GetStation())
	if err != nil {
		return nil, err
	}
	s2n := NewSend2Node(sta, c.unpackState(req.State), c.env)
	err = s2n.Send2Esp()

	if err != nil {
		return &pb.ResponseSync{
			Response: pb.ResponseSync_FAILED,
		}, err
	}
	return &pb.ResponseSync{Response: pb.ResponseSync_SUCCESS}, nil
}

func (c *ControlServer) unpackStations(req *pb.Stations) (*pkg.StationDetail, error) {
	s, ok := pb.Stations_StationId_name[int32(req.GetStationId())]
	if !ok {
		return nil, fmt.Errorf("station: %s is not found in proto file\n", req.String())
	}
	sta, err := c.Stations.SearchStation(s)
	if err != nil {
		return nil, fmt.Errorf("station %s is not found in yaml file\n", s)
	}
	return sta, nil
}
func (c *ControlServer) unpackState(state pb.RequestSync_State) string {
	return state.String()
}
