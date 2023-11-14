package _1334

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findTheCity(t *testing.T) {
	type args struct {
		n                 int
		edges             [][]int
		distanceThreshold int
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
			if got := findTheCity(tt.args.n, tt.args.edges, tt.args.distanceThreshold); got != tt.want {
				t.Errorf("findTheCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
