package _2327

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_peopleAwareOfSecret(t *testing.T) {
	type args struct {
		n      int
		delay  int
		forget int
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
			if got := peopleAwareOfSecret(tt.args.n, tt.args.delay, tt.args.forget); got != tt.want {
				t.Errorf("peopleAwareOfSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
