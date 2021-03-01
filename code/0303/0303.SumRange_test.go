package _303

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		sum []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NumArray{
				sum: tt.fields.sum,
			}
			if got := n.SumRange(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
