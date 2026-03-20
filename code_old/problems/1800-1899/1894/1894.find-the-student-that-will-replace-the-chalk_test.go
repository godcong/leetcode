package _1894

import (
	"testing"
)

func Test_chalkReplacer(t *testing.T) {
	type args struct {
		chalk []int
		k     int
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
			if got := chalkReplacer(tt.args.chalk, tt.args.k); got != tt.want {
				t.Errorf("chalkReplacer() = %v, want %v", got, tt.want)
			}
		})
	}
}
