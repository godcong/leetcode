package _1146

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want SnapshotArray
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnapshotArray_Set(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *SnapshotArray
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Set(tt.args.index, tt.args.val)
		})
	}
}

func TestSnapshotArray_Snap(t *testing.T) {
	tests := []struct {
		name string
		this *SnapshotArray
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Snap(); got != tt.want {
				t.Errorf("SnapshotArray.Snap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnapshotArray_Get(t *testing.T) {
	type args struct {
		index   int
		snap_id int
	}
	tests := []struct {
		name string
		this *SnapshotArray
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.index, tt.args.snap_id); got != tt.want {
				t.Errorf("SnapshotArray.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
