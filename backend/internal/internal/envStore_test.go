package internal

import (
	"testing"
)

func TestPort_String(t *testing.T) {
	tests := []struct {
		name string
		p    Port
		want string
	}{
		{
			name: "port 0",
			p:    0,
			want: "0",
		},
		{
			name: "port 1",
			p:    1,
			want: "1",
		},
		{
			name: "port 65535",
			p:    65535,
			want: "65535",
		},
		{
			name: "port 65536",
			p:    65536,
			want: "65536",
		},
		{
			name: "port 65537",
			p:    65537,
			want: "65537",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPort_Unmarshal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		p       Port
		args    args
		wantErr bool
	}{
		{
			name:    "port -1",
			args:    args{s: "-1"},
			wantErr: true,
		},
		{
			name:    "port 0",
			args:    args{s: "0"},
			wantErr: true,
		},
		{
			name:    "port 1",
			args:    args{s: "1"},
			wantErr: false,
		},
		{
			name:    "port 65535",
			args:    args{s: "65535"},
			wantErr: false,
		},
		{
			name:    "port 65536",
			args:    args{s: "65536"},
			wantErr: true,
		},
		{
			name:    "port is not numeric",
			args:    args{s: "port"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Unmarshal(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
