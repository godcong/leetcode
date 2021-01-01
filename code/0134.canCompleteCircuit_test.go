package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
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
				gas:  leetcode.strToIntArray("[1,2,3,4,5]"),
				cost: leetcode.strToIntArray("[3,4,5,1,2]"),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				gas:  leetcode.strToIntArray("[2,3,4]"),
				cost: leetcode.strToIntArray("[3,4,3]"),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
