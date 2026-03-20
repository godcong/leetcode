package _2928

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		n     int
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
			if got := distributeCandies(tt.args.n, tt.args.limit); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
