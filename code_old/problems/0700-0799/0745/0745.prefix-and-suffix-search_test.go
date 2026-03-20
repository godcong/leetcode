package _0745

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want WordFilter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordFilter_F(t *testing.T) {
	type args struct {
		pref string
		suff string
	}
	tests := []struct {
		name string
		this *WordFilter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.F(tt.args.pref, tt.args.suff); got != tt.want {
				t.Errorf("WordFilter.F() = %v, want %v", got, tt.want)
			}
		})
	}
}
