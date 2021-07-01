package _LCP_07

import (
	"testing"
)

func Test_numWays(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
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
			if got := numWays(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
