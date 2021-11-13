package msg2Esp

//func TestNewSend2Node(t *testing.T) {
//	type args struct {
//		sta   *station2espIp.StationDetail
//		state string
//		e     *internal.Env
//	}
//	tests := []struct {
//		name string
//		args args
//		want *send2node
//	}{
//		{
//			name: "",
//			args: args{
//				sta: &station2espIp.StationDetail{
//					Pin: 0,
//				},
//				state: "ON",
//				e:     nil,
//			},
//			want: &send2node{
//				Station:     &station2espIp.StationDetail{Pin: 0},
//				Environment: nil,
//				sendData: &sendData{
//					State: "ON",
//					Pin:   0,
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewSend2Node(tt.args.sta, tt.args.state, tt.args.e); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewSend2Node() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
