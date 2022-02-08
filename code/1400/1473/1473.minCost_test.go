package _1473

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				houses: nil,
				cost:   nil,
				m:      0,
				n:      0,
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.houses, tt.args.cost, tt.args.m, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
