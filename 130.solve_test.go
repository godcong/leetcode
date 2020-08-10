package leetcode

import (
	"bytes"
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
			for i := range tt.args.board {
				if bytes.Compare(tt.args.board[i], tt.want[i]) != 0 {
					t.Errorf("solve(%d) = %s, want %s", i, tt.args.board[i], tt.want[i])
				}
			}
		})
	}
}
