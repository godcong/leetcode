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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0188() = %v, want %v", got, tt.want)
			}
		})
	}
}
