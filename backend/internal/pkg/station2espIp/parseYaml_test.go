package station2espIp

import (
	"reflect"
	"testing"
)

func TestStations_SearchStation(t *testing.T) {
	type fields struct {
		Stations []Station
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *StationDetail
		wantErr bool
	}{
		{
			name: "exist station",
			fields: fields{Stations: []Station{
				{Station: StationDetail{Name: "CHOFU"}},
				{Station: StationDetail{Name: "FUCHU"}},
			}},
			args:    args{name: "CHOFU"},
			want:    &StationDetail{Name: "CHOFU"},
			wantErr: false,
		},
		{
			name: "small capital",
			fields: fields{Stations: []Station{
				{Station: StationDetail{Name: "CHOFU"}},
				{Station: StationDetail{Name: "FUCHU"}},
			}},
			args:    args{name: "chofu"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "station not exist",
			fields: fields{Stations: []Station{
				{Station: StationDetail{Name: "CHOFU"}},
				{Station: StationDetail{Name: "FUCHU"}},
			}},
			args:    args{name: "SHINJUKU"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stations{
				Stations: tt.fields.Stations,
			}
			got, err := s.SearchStation(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchStation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchStation() got = %v, want %v", got, tt.want)
			}
		})
	}
}
