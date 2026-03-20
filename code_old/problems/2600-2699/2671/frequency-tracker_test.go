package _2671

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want FrequencyTracker
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

func TestFrequencyTracker_Add(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		this *FrequencyTracker
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.number)
		})
	}
}

func TestFrequencyTracker_DeleteOne(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		this *FrequencyTracker
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.DeleteOne(tt.args.number)
		})
	}
}

func TestFrequencyTracker_HasFrequency(t *testing.T) {
	type args struct {
		frequency int
	}
	tests := []struct {
		name string
		this *FrequencyTracker
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.HasFrequency(tt.args.frequency); got != tt.want {
				t.Errorf("FrequencyTracker.HasFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
