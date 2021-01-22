package code

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: strToIntArray("[1,2,0,0]"),
				K: 34,
			},
			wantAns: strToIntArray("[1,2,3,4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := addToArrayForm(tt.args.A, tt.args.K); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("addToArrayForm() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
