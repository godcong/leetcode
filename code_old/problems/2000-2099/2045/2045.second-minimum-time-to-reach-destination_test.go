package _2045

import (
	"testing"
)

func Test_secondMinimum(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		time   int
		change int
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
			if got := secondMinimum(tt.args.n, tt.args.edges, tt.args.time, tt.args.change); got != tt.want {
				t.Errorf("secondMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
