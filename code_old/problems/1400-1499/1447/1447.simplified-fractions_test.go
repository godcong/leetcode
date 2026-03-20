package _1447

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
