package leetcode

import "testing"

func Test_maxProfit0188(t *testing.T) {
	type args struct {
		k      int
		prices []int
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
				k:      2,
				prices: strToIntArray("[2,4,1]"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0188() = %v, want %v", got, tt.want)
			}
		})
	}
}
