package _3606

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_validateCoupons(t *testing.T) {
	type args struct {
		code         []string
		businessLine []string
		isActive     []bool
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
			if got := validateCoupons(tt.args.code, tt.args.businessLine, tt.args.isActive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
