package code

import "testing"

func Test_videoStitching(t *testing.T) {
	type args struct {
		clips [][]int
		T     int
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
				clips: [][]int{
					{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9},
				},
				T: 10,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				clips: [][]int{
					{0, 1}, {1, 2},
				},
				T: 5,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				clips: [][]int{
					{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9},
				},
				T: 9,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				clips: [][]int{
					{0, 4}, {2, 8},
				},
				T: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := videoStitching(tt.args.clips, tt.args.T); got != tt.want {
				t.Errorf("videoStitching() = %v, want %v", got, tt.want)
			}
		})
	}
}
