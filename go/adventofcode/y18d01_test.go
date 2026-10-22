package adventofcode

import "testing"

var y18d01Tests []string = []string{
	"+1\n+1\n+1\n",
	"+1\n+1\n-2\n",
	"-1\n-2\n-3\n",
	"+1\n-1\n",
	"+3\n+3\n+4\n-2\n-4\n",
	"-6\n+3\n+8\n+5\n-6\n",
	"+7\n+7\n-2\n-7\n-4\n",
}

func Test_y18d01(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{name: "Test_y18d01_1", args: args{input: y18d01Tests[0], part: 1}, wantAnswer: 3},
		{name: "Test_y18d01_2", args: args{input: y18d01Tests[1], part: 1}, wantAnswer: 0},
		{name: "Test_y18d01_3", args: args{input: y18d01Tests[2], part: 1}, wantAnswer: -6},
		{name: "Test_y18d01_4", args: args{input: y18d01Tests[3], part: 2}, wantAnswer: 0},
		{name: "Test_y18d01_5", args: args{input: y18d01Tests[4], part: 2}, wantAnswer: 10},
		{name: "Test_y18d01_6", args: args{input: y18d01Tests[5], part: 2}, wantAnswer: 5},
		{name: "Test_y18d01_7", args: args{input: y18d01Tests[6], part: 2}, wantAnswer: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := y18d01(tt.args.input, tt.args.part); gotAnswer != tt.wantAnswer {
				t.Errorf("y18d01() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
