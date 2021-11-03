package envStore

import (
	"testing"
)

func Test_hostnamePort_Unmarshal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		t       hostnamePort
		args    args
		wantErr bool
	}{
		{
			name:    "localhost",
			t:       "",
			args:    args{"127.0.0.1:12345"},
			wantErr: false,
		},
		{
			name:    "localhost_noPort",
			t:       "",
			args:    args{"127.0.0.1"},
			wantErr: true,
		},
		{
			name:    "0.0.0.0",
			t:       "",
			args:    args{"0.0.0.0:12345"},
			wantErr: false,
		},
		{
			name:    "example.com",
			t:       "",
			args:    args{"example.com:12345"},
			wantErr: false,
		},
		{
			name:    "range err Port",
			t:       "",
			args:    args{"example.com:65536"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.Unmarshal(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
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
			name:    "negative port",
			p:       0,
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "zero port",
			p:       0,
			args:    args{"0"},
			wantErr: true,
		},
		{
			name:    "port 1",
			p:       0,
			args:    args{"1"},
			wantErr: false,
		},
		{
			name:    "port 65535",
			p:       0,
			args:    args{"65535"},
			wantErr: false,
		},
		{
			name:    "port 65536",
			p:       0,
			args:    args{"65536"},
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
