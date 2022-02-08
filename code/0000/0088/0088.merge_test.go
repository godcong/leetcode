package _0088

import (
	. "github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
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
				nums1: StrToIntArray("[1,2,3,0,0,0]"),
				m:     3,
				nums2: StrToIntArray("[2,5,6]"),
				n:     3,
			},
			want: StrToIntArray("[1,2,2,3,5,6]"),
		},
		{
			name: "",
			args: args{
				nums1: StrToIntArray("[1]"),
				m:     1,
				nums2: StrToIntArray("[]"),
				n:     0,
			},
			want: StrToIntArray("[1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
