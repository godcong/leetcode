package _1711

import (
	"testing"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		deliciousness []int
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
			if got := countPairs(tt.args.deliciousness); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
