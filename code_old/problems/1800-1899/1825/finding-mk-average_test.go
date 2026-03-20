package _1825

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		m int
		k int
	}
	tests := []struct {
		name string
		args args
		want MKAverage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.m, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMKAverage_AddElement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *MKAverage
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddElement(tt.args.num)
		})
	}
}

func TestMKAverage_CalculateMKAverage(t *testing.T) {
	tests := []struct {
		name string
		this *MKAverage
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.CalculateMKAverage(); got != tt.want {
				t.Errorf("MKAverage.CalculateMKAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
