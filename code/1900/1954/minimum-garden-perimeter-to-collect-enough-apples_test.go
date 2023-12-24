package _1954

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumPerimeter(t *testing.T) {
	type args struct {
		neededApples int64
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
			if got := minimumPerimeter(tt.args.neededApples); got != tt.want {
				t.Errorf("minimumPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
