package _1736

import (
	"testing"
)

func Test_maximumTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTime(tt.args.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
