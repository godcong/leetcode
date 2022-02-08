package old

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
				[]byte("XXOXXX"),
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
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("OXOOXX"),
					[]byte("OXXXOX"),
					[]byte("XOOXOO"),
					[]byte("XOXXXX"),
					[]byte("OOXOXX"),
					[]byte("XXOOOO"),
				},
			},
			want: [][]byte{
				[]byte("OXOOXX"),
				[]byte("OXXXOX"),
				[]byte("XOOXOO"),
				[]byte("XOXXXX"),
				[]byte("OOXOXX"),
				[]byte("XXOOOO"),
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					[]byte("OXOOXX"),
					[]byte("XOXXOO"),
					[]byte("XOOOOX"),
					[]byte("XOXXXX"),
					[]byte("XOXOXX"),
					[]byte("XXOOOO"),
				},
			},
			want: [][]byte{
				[]byte("OXOOXX"),
				[]byte("XOXXOO"),
				[]byte("XOOOOX"),
				[]byte("XOXXXX"),
				[]byte("XOXOXX"),
				[]byte("XXOOOO"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			t.Logf("board:%s\n", tt.args.board)
			for i := range tt.args.board {
				if bytes.Compare(tt.args.board[i], tt.want[i]) != 0 {
					t.Errorf("solve(%d) = %s, want %s", i, tt.args.board[i], tt.want[i])
				}
			}
		})
	}
}
