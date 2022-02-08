package old

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				s: "25525511135",
			},
			want: []string{
				"255.255.11.135", "255.255.111.35",
			},
		},
		{
			name: "",
			args: args{
				s: "1921680102",
			},
			want: []string{

				"19.216.80.102", "192.16.80.102", "192.168.0.102",
			},
		},
		{
			name: "",
			args: args{
				s: "1921681112",
			},
			want: []string{
				"19.216.81.112", "192.16.81.112", "192.168.1.112", "192.168.11.12", "192.168.111.2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
