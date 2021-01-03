package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
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
				arr: strToIntArray("[0,1,2,3,4,5,6,7,8]"),
			},
			want: strToIntArray("[0,1,2,4,8,3,5,6,7]"),
		},
		{
			name: "",
			args: args{
				arr: strToIntArray("[1024,512,256,128,64,32,16,8,4,2,1]"),
			},
			want: strToIntArray("[1,2,4,8,16,32,64,128,256,512,1024]"),
		},
		{
			name: "",
			args: args{
				arr: strToIntArray("[10000,10000]"),
			},
			want: strToIntArray("[10000,10000]"),
		},
		{
			name: "",
			args: args{
				arr: strToIntArray("[2,3,5,7,11,13,17,19]"),
			},
			want: strToIntArray("[2,3,5,17,7,11,13,19]"),
		},
		{
			name: "",
			args: args{
				arr: strToIntArray("[10,100,1000,10000]"),
			},
			want: strToIntArray("[10,100,10000,1000]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
