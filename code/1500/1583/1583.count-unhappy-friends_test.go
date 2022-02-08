package _1583

import (
	"testing"
)

func Test_unhappyFriends(t *testing.T) {
	type args struct {
		n           int
		preferences [][]int
		pairs       [][]int
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
			if got := unhappyFriends(tt.args.n, tt.args.preferences, tt.args.pairs); got != tt.want {
				t.Errorf("unhappyFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
