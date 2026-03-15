package _1622

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Fancy
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

func TestFancy_Append(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *Fancy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Append(tt.args.val)
		})
	}
}

func TestFancy_AddAll(t *testing.T) {
	type args struct {
		inc int
	}
	tests := []struct {
		name string
		this *Fancy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAll(tt.args.inc)
		})
	}
}

func TestFancy_MultAll(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		this *Fancy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.MultAll(tt.args.m)
		})
	}
}

func TestFancy_GetIndex(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		this *Fancy
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetIndex(tt.args.idx); got != tt.want {
				t.Errorf("Fancy.GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
