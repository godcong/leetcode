package leetcode

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				[]int{2, 5, 1, 3, 4, 7}, 3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "",
			args: args{
				[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4,
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "",
			args: args{
				[]int{1, 1, 2, 2}, 2,
			},
			want: []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
