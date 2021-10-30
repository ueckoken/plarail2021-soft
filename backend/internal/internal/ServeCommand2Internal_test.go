package internal

import (
	"reflect"
	"testing"
	"ueckoken/plarail2021-soft-internal/pkg"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

func TestControlServer_unpackState(t *testing.T) {
	type fields struct {
		UnimplementedControlServer pb.UnimplementedControlServer
		env                        *Env
		Stations                   *pkg.Stations
	}
	type args struct {
		state pb.RequestSync_State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "state is on",
			args: args{state: pb.RequestSync_ON},
			want: "ON",
		},
		{
			name: "state is off",
			args: args{state: pb.RequestSync_OFF},
			want: "OFF",
		},
		{
			name: "state is unknown",
			args: args{state: pb.RequestSync_UNKNOWN},
			want: "UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ControlServer{
				UnimplementedControlServer: tt.fields.UnimplementedControlServer,
				env:                        tt.fields.env,
				Stations:                   tt.fields.Stations,
			}
			if got := c.unpackState(tt.args.state); got != tt.want {
				t.Errorf("unpackState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestControlServer_unpackStations(t *testing.T) {
	type fields struct {
		UnimplementedControlServer pb.UnimplementedControlServer
		env                        *Env
		Stations                   *pkg.Stations
	}
	type args struct {
		req *pb.Stations
	}
	f := fields{
		Stations: &pkg.Stations{Stations: []pkg.Station{
			{
				pkg.StationDetail{
					Name:    "chofu_b1",
					Address: "TEST_ADDR",
					Pin:     1,
				},
			}, {
				pkg.StationDetail{
					Name:    "chofu_b2",
					Address: "TEST_ADDR",
					Pin:     2,
				},
			},
			{
				pkg.StationDetail{
					Name:    "TOKYO",
					Address: "TEST_ADDR",
					Pin:     1,
				},
			}, {
				pkg.StationDetail{
					Name:    "chofu_b2",
					Address: "TEST_ADDR",
					Pin:     2,
				},
			},
		}}}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pkg.StationDetail
		wantErr bool
	}{
		{
			name:   "station exist",
			fields: f,
			args:   args{req: &pb.Stations{StationId: pb.Stations_chofu_b1}},
			want: &pkg.StationDetail{
				Name:    "chofu_b1",
				Address: "TEST_ADDR",
				Pin:     1,
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: f,
			args:   args{req: &pb.Stations{StationId: pb.Stations_chofu_b1}},
			want: &pkg.StationDetail{
				Name:    "chofu_b1",
				Address: "TEST_ADDR",
				Pin:     1,
			},
			wantErr: false,
		}, {
			name:    "station not define in yaml",
			fields:  f,
			args:    args{req: &pb.Stations{StationId: pb.Stations_sasazuka_b1}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ControlServer{
				UnimplementedControlServer: tt.fields.UnimplementedControlServer,
				env:                        tt.fields.env,
				Stations:                   tt.fields.Stations,
			}
			got, err := c.unpackStations(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackStations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackStations() got = %v, want %v", got, tt.want)
			}
		})
	}
}