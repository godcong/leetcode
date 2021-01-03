package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
				nums1: strToIntArray("[1,2,2,1]"),
				nums2: strToIntArray("[2,2]"),
			},
			want: strToIntArray("[2]"),
		},
		{
			name: "",
			args: args{
				nums1: strToIntArray("[4,9,5]"),
				nums2: strToIntArray("[9,4,9,8,4]"),
			},
			want: strToIntArray("[9,4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
