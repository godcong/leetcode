package _2999

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfPowerfulInt(t *testing.T) {
	type args struct {
		start  int64
		finish int64
		limit  int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPowerfulInt(tt.args.start, tt.args.finish, tt.args.limit, tt.args.s); got != tt.want {
				t.Errorf("numberOfPowerfulInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
