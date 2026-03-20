package _0810

import (
	"testing"
)

func Test_xorGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorGame(tt.args.nums); got != tt.want {
				t.Errorf("xorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
