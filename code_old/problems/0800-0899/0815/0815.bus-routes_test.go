package _0815

import (
	"testing"
)

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
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
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
