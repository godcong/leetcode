package _3484

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		rows int
	}
	tests := []struct {
		name string
		args args
		want Spreadsheet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpreadsheet_SetCell(t *testing.T) {
	type args struct {
		cell  string
		value int
	}
	tests := []struct {
		name string
		this *Spreadsheet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.SetCell(tt.args.cell, tt.args.value)
		})
	}
}

func TestSpreadsheet_ResetCell(t *testing.T) {
	type args struct {
		cell string
	}
	tests := []struct {
		name string
		this *Spreadsheet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.ResetCell(tt.args.cell)
		})
	}
}

func TestSpreadsheet_GetValue(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		this *Spreadsheet
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetValue(tt.args.formula); got != tt.want {
				t.Errorf("Spreadsheet.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
