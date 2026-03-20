package _1705

import (
	"testing"
)

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
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
			if got := eatenApples(tt.args.apples, tt.args.days); got != tt.want {
				t.Errorf("eatenApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
