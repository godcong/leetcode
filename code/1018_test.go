package code

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: strToIntArray("[0,1,1]"),
			},
			want: []bool{
				true, false, false,
			},
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,1,1]"),
			},
			want: []bool{
				false, false, false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
