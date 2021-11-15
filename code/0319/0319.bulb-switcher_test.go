package _0319

import (
	"testing"
)

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
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
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
