package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *leetcode.Node
	}
	tests := []struct {
		name string
		args args
		want *leetcode.Node
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect117(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
