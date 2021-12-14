package _0630

import (
	"testing"
)

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
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
			if got := scheduleCourse(tt.args.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
