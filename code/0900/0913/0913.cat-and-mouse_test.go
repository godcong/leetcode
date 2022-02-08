package _0913

import (
	"testing"
)

func Test_catMouseGame(t *testing.T) {
	type args struct {
		graph [][]int
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
			if got := catMouseGame(tt.args.graph); got != tt.want {
				t.Errorf("catMouseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
