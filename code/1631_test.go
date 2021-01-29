package code

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
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
				heights: strToIntArrArray("[[1,2,2],[3,8,2],[5,3,5]]"),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				heights: strToIntArrArray("[[1,2,3],[3,8,4],[5,3,5]]"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
