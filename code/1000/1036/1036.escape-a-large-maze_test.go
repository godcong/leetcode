package _1036

import (
	"testing"
)

func Test_isEscapePossible(t *testing.T) {
	type args struct {
		blocked [][]int
		source  []int
		target  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.args.blocked, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("isEscapePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
