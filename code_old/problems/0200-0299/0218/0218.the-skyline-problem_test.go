package _0218

import (
	"reflect"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	type args struct {
		buildings [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getSkyline(tt.args.buildings); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getSkyline() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
