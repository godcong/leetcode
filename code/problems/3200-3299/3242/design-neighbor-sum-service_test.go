package _3242

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want NeighborSum
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighborSum_AdjacentSum(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *NeighborSum
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.AdjacentSum(tt.args.value); got != tt.want {
				t.Errorf("NeighborSum.AdjacentSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeighborSum_DiagonalSum(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *NeighborSum
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.DiagonalSum(tt.args.value); got != tt.want {
				t.Errorf("NeighborSum.DiagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
