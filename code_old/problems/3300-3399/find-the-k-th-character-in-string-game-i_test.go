package _3304

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
