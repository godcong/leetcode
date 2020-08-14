package leetcode

import "testing"

func Test_removeBoxes(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				boxes: []int{
					1, 1,
				},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 2,
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 1, 2,
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 2, 1,
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 1, 1,
				},
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 2, 1, 2, 1,
				},
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 3, 2, 2, 2, 3, 4, 3, 1,
				},
			},
			want: 23,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 3, 2, 2, 2, 3, 4, 3, 1,
				},
			},
			want: 23,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					1, 3, 2, 2, 4, 2, 3, 4, 3, 1,
				},
			},
			want: 24,
		},
		{
			name: "",
			args: args{
				boxes: []int{
					3, 8, 8, 5, 5, 3, 9, 2, 4, 4, 6, 5, 8, 4, 8, 6, 9, 6, 2, 8, 6, 4, 1, 9, 5, 3, 10, 5, 3, 3, 9, 8, 8, 6, 5, 3, 7, 4, 9, 6, 3, 9, 4, 3, 5, 10, 7, 6, 10, 7,
				},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
