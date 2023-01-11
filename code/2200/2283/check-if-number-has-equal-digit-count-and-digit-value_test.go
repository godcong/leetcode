package _2283

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_digitCount(t *testing.T) {
	type args struct {
		num string
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
			if got := digitCount(tt.args.num); got != tt.want {
				t.Errorf("digitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
