package old

import "testing"

func Test_numSimilarGroups(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				strs: strToStringArray("[\"tars\",\"rats\",\"arts\",\"star\"]"),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				strs: strToStringArray("[\"omv\",\"ovm\"]"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSimilarGroups(tt.args.strs); got != tt.want {
				t.Errorf("numSimilarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
