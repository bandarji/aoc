package adventofcode

import "testing"

func Test_y17d17(t *testing.T) {
	type args struct {
		input         string
		valueToInsert int
		valueToFind   int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name: "test 1",
			args: args{
				input:         y17d17StartValue,
				valueToInsert: 2017,
				valueToFind:   2017,
			},
			wantAnswer: 1670,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y17d17(tt.args.input, tt.args.valueToInsert, tt.args.valueToFind); gotAnswer != tt.wantAnswer {
				t.Errorf("y17d17() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
