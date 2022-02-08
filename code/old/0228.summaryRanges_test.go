package old

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: strToIntArray("[0,1,2,4,5,7]"),
			},
			want: strToStringArray("[\"0->2\",\"4->5\",\"7\"]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[0,2,3,4,6,8,9]"),
			},
			want: strToStringArray("[\"0\",\"2->4\",\"6\",\"8->9\"]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[]"),
			},
			want: strToStringArray("[]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray(" [-1]"),
			},
			want: strToStringArray(" [-1]"),
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[0]"),
			},
			want: strToStringArray("[0]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
