package code

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
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
				arr1: strToIntArray("[2,3,1,3,2,4,6,7,9,2,19]"),
				arr2: strToIntArray("[2,1,4,3,9,6]"),
			},
			want: strToIntArray("[2,2,2,1,4,3,3,9,6,7,19]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
