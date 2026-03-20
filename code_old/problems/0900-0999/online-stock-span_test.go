package _0901

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want StockSpanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockSpanner_Next(t *testing.T) {
	type args struct {
		price int
	}
	tests := []struct {
		name string
		this *StockSpanner
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Next(tt.args.price); got != tt.want {
				t.Errorf("StockSpanner.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
