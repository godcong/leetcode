package old

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
		{
			name: "",
			args: args{
				A: strToIntArray("[2,7,4]"),
				K: 181,
			},
			wantAns: strToIntArray("[4,5,5]"),
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[2,1,5]"),
				K: 806,
			},
			wantAns: strToIntArray("[1,0,2,1]"),
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[9,9,9,9,9,9,9,9,9,9]"),
				K: 1,
			},
			wantAns: strToIntArray("[1,0,0,0,0,0,0,0,0,0,0]"),
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
