package leetcode

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				commands:  nil,
				obstacles: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}
