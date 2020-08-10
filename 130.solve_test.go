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
					[]byte("O"),
				},
			},
			want: [][]byte{
				[]byte("O"),
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("OO"),
					[]byte("OO"),
				},
			},
			want: [][]byte{
				[]byte("OO"),
				[]byte("OO"),
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("OOO"),
					[]byte("OOO"),
					[]byte("OOO"),
				},
			},
			want: [][]byte{
				[]byte("OOO"),
				[]byte("OOO"),
				[]byte("OOO"),
			},
		},
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
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("XXXXX"),
					[]byte("XXOOX"),
					[]byte("XXXOX"),
					[]byte("XXOXX"),
					[]byte("XXOXX"),
				},
			},
			want: [][]byte{
				[]byte("XXXXX"),
				[]byte("XXXXX"),
				[]byte("XXXXX"),
				[]byte("XXOXX"),
				[]byte("XXOXX"),
			},
		},

		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("XOXOXO"),
					[]byte("OXOXOX"),
					[]byte("XOXOXO"),
					[]byte("OXOXOX"),
				},
			},
			want: [][]byte{
				[]byte("XOXOXO"),
				[]byte("OXXXXX"),
				[]byte("XXXXXO"),
				[]byte("OXOXOX"),
				[]byte("XXOXX"),
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("XXXXOX"),
					[]byte("OXXOOX"),
					[]byte("XOXOOO"),
					[]byte("XOOOXO"),
					[]byte("OOXXOX"),
					[]byte("XOXOXX"),
				},
			},
			want: [][]byte{
				[]byte("XXXXOX"),
				[]byte("OXXOOX"),
				[]byte("XOXOOO"),
				[]byte("XOOOXO"),
				[]byte("OOXXXX"),
				[]byte("XOXOXX"),
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
