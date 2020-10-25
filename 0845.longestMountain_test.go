package leetcode

import "testing"

func Test_longestMountain(t *testing.T) {
	type args struct {
		A []int
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
				A: []int{
					2, 1, 4, 7, 3, 2, 5,
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				A: []int{
					2, 2, 2,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.A); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
