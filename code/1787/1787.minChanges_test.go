package _1787

import (
	"testing"
)

func Test_minChanges(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := minChanges(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
