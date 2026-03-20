package _SwordRefers_Offer_II_041

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want MovingAverage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovingAverage_Next(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MovingAverage
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Next(tt.args.val); got != tt.want {
				t.Errorf("MovingAverage.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
