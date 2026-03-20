package _0752

import (
	"testing"
)

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
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
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
