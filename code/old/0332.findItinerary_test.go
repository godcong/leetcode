package old

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				tickets: [][]string{
					{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
				},
			},
			want: []string{
				"JFK", "MUC", "LHR", "SFO", "SJC",
			},
		},
		{
			name: "",
			args: args{
				tickets: [][]string{
					{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
				},
			},
			want: []string{
				"JFK", "ATL", "JFK", "SFO", "ATL", "SFO",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
