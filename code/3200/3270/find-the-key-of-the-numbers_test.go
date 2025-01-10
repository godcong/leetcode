package _3270

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_generateKey(t *testing.T) {
	type args struct {
		num1 int
		num2 int
		num3 int
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
			if got := generateKey(tt.args.num1, tt.args.num2, tt.args.num3); got != tt.want {
				t.Errorf("generateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
