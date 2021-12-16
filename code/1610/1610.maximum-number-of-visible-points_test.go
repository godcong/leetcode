package _1610

import (
	"testing"
)

func Test_visiblePoints(t *testing.T) {
	type args struct {
		points   [][]int
		angle    int
		location []int
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
			if got := visiblePoints(tt.args.points, tt.args.angle, tt.args.location); got != tt.want {
				t.Errorf("visiblePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
