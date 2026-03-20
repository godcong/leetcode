package _1269

import (
	"testing"
)

func Test_numWays(t *testing.T) {
	type args struct {
		steps  int
		arrLen int
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
			if got := numWays(tt.args.steps, tt.args.arrLen); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
