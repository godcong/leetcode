package _0664

import (
	"testing"
)

func Test_strangePrinter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "aaabbb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
