package leetcode

import "testing"

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				heights: []int{
					1, 1, 4, 2, 1, 3,
				},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				heights: []int{
					5, 1, 2, 3, 4,
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				heights: []int{
					1, 2, 3, 4, 5,
				},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				heights: []int{
					2, 1, 2, 1, 1, 2, 2, 1,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.args.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
