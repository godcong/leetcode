package _1105

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minHeightShelves(t *testing.T) {
	type args struct {
		books      [][]int
		shelfWidth int
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
			if got := minHeightShelves(tt.args.books, tt.args.shelfWidth); got != tt.want {
				t.Errorf("minHeightShelves() = %v, want %v", got, tt.want)
			}
		})
	}
}
