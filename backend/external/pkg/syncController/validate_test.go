package syncController

import (
	"testing"
	"ueckoken/plarail2021-soft-external/pkg/servo"
)

const (
	CHOFU_B1 = 1
	CHOFU_B2 = 2
	CHOFU_B3 = 3
	CHOFU_S1 = 4
	CHOFU_S2 = 5
	CHOFU_S3 = 6
	CHOFU_S4 = 7
)
const (
	ON  = 1
	OFF = 2
)

func TestValidator_Validate(t *testing.T) {
	type fields struct {
		Stations []Station
	}
	type args struct {
		u  StationState
		ss []StationState
	}
	rules := []Station{{
		EachStation{
			Name:   "chofu_kudari",
			Points: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
			Rules: []Rule{
				{
					On:  nil,
					Off: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
				},
				{
					On:  []string{"chofu_s1"},
					Off: nil,
				},
				{
					On:  []string{"chofu_s2"},
					Off: nil,
				},
			},
		},
	},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "ルールの1つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: CHOFU_B1, State: OFF}},
				ss: []StationState{
					{servo.StationState{StationID: CHOFU_B1, State: ON}},  // chofu_b1,on
					{servo.StationState{StationID: CHOFU_B2, State: OFF}}, // chofu_b2,off
					{servo.StationState{StationID: CHOFU_S1, State: OFF}}, // chofu_s1,off
					{servo.StationState{StationID: CHOFU_S2, State: OFF}}, // chofu_s2,off
				},
			},
			wantErr: false,
		},
		{
			name:   "ルールの2つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: CHOFU_S1, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: CHOFU_B1, State: ON}},
					{servo.StationState{StationID: CHOFU_B2, State: ON}},
					{servo.StationState{StationID: CHOFU_S1, State: OFF}},
					{servo.StationState{StationID: CHOFU_S2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:   "ルールの3つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: CHOFU_S2, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: CHOFU_B1, State: ON}},
					{servo.StationState{StationID: CHOFU_B2, State: ON}},
					{servo.StationState{StationID: CHOFU_S1, State: OFF}},
					{servo.StationState{StationID: CHOFU_S2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:   "2,3つ目のルールはONが複数あっても良い",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: CHOFU_S2, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: CHOFU_B1, State: ON}},
					{servo.StationState{StationID: CHOFU_B2, State: ON}},
					{servo.StationState{StationID: CHOFU_S1, State: ON}},
					{servo.StationState{StationID: CHOFU_S2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:    "バリデートの対象外",
			fields:  fields{Stations: rules},
			args:    args{u: StationState{servo.StationState{StationID: 10, State: OFF}}},
			wantErr: false,
		},
		{
			name:   "3つ目のルール違反",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: CHOFU_B1, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: CHOFU_B1, State: OFF}},
					{servo.StationState{StationID: CHOFU_B2, State: OFF}},
					{servo.StationState{StationID: CHOFU_S1, State: OFF}},
					{servo.StationState{StationID: CHOFU_S2, State: OFF}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				Stations: tt.fields.Stations,
			}
			if err := v.Validate(tt.args.u, tt.args.ss); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
