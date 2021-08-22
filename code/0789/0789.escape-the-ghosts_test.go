package _0789

import (
	"testing"
)

func Test_escapeGhosts(t *testing.T) {
	type args struct {
		ghosts [][]int
		target []int
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
			if got := escapeGhosts(tt.args.ghosts, tt.args.target); got != tt.want {
				t.Errorf("escapeGhosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
