package _3129

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfStableArrays(t *testing.T) {
	type args struct {
		zero  int
		one   int
		limit int
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
			if got := numberOfStableArrays(tt.args.zero, tt.args.one, tt.args.limit); got != tt.want {
				t.Errorf("numberOfStableArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
