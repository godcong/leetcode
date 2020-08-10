package leetcode

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("XXXX"),
					[]byte("XOOX"),
					[]byte("XXOX"),
					[]byte("XOXX"),
				},
			},
			want: [][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XOXX"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solve() = %s, want %s", tt.args.board, tt.want)
			}
		})
	}
}
