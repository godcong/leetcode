package code

import (
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	type args struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
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
				n:           8,
				m:           2,
				group:       strToIntArray("[-1,-1,1,0,0,1,0,-1]"),
				beforeItems: strToIntArrArray("[[],[6],[5],[6],[3,6],[],[],[]]"),
			},
			want: strToIntArray("[6,3,4,1,5,2,0,7]"),
		},
		{
			name: "",
			args: args{
				n:           8,
				m:           2,
				group:       strToIntArray("[-1,-1,1,0,0,1,0,-1]"),
				beforeItems: strToIntArrArray("[[],[6],[5],[6],[3],[],[4],[]]"),
			},
			want: strToIntArray("[]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.n, tt.args.m, tt.args.group, tt.args.beforeItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
