package leetcode

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
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
				matrix: strToByteArrayArray("[[\"1\",\"0\",\"1\",\"0\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"0\"]]"),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				matrix: strToByteArrayArray("[]"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				matrix: strToByteArrayArray("[[\"0\"]]"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				matrix: strToByteArrayArray("[[\"1\"]]"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				matrix: strToByteArrayArray("[[\"0\",\"0\"]]"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
