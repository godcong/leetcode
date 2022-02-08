package old

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				intervals:   strToIntArrArray("[[1,3],[6,9]]"),
				newInterval: strToIntArray("[2,5]"),
			},
			want: strToIntArrArray("[[1,5],[6,9]]"),
		},
		{
			name: "",
			args: args{
				intervals:   strToIntArrArray("[[1,2],[3,5],[6,7],[8,10],[12,16]]"),
				newInterval: strToIntArray("[4,8]"),
			},
			want: strToIntArrArray("[[1,2],[3,10],[12,16]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
