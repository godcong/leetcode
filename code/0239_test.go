package code

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: strToIntArray("[1,3,-1,-3,5,3,6,7]"),
				k:    3,
			},
			want: strToIntArray("[3,3,5,5,6,7]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[1]"),
				k:    1,
			},
			want: strToIntArray("[1]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[1,-1]"),
				k:    1,
			},
			want: strToIntArray("[1,-1]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[9,11]"),
				k:    2,
			},
			want: strToIntArray("[11]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
