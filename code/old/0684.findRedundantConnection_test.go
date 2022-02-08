package old

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
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
				edges: strToIntArrArray("[[1,2], [1,3], [2,3]]"),
			},
			want: strToIntArray("[2,3]"),
		},
		{
			name: "",
			args: args{
				edges: strToIntArrArray("[[1,2], [2,3], [3,4], [1,4], [1,5]]"),
			},
			want: strToIntArray("[1,4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
