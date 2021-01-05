package code

import (
	"testing"
)

func Test_countRangeSum(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
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
				nums:  strToIntArray("[-2,5,-1]"),
				lower: -2,
				upper: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSum(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
