package _1720

import (
	"github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		encoded []int
		first   int
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
				encoded: common.StrToIntArray("[1,2,3]"),
				first:   1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := decode(tt.args.encoded, tt.args.first); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("decode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
