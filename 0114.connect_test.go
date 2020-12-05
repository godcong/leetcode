package leetcode

import (
	"reflect"
	"testing"
)

func Test_connect114(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
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
			if got := connect114(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connedct114() = %v, want %v", got, tt.want)
			}
		})
	}
}
