package _0405

import (
	"testing"
)

func Test_toHex(t *testing.T) {
	type args struct {
		num int
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
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
