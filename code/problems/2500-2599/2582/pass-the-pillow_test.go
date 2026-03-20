package _2582

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_passThePillow(t *testing.T) {
	type args struct {
		n    int
		time int
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
			if got := passThePillow(tt.args.n, tt.args.time); got != tt.want {
				t.Errorf("passThePillow() = %v, want %v", got, tt.want)
			}
		})
	}
}
