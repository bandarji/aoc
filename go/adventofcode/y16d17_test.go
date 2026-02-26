package adventofcode

import "testing"

func Test_y16d17MD5BFS(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 'ihgpwlah'", args{input: "ihgpwlah", part: 1}, "DDRRRD"},
		{"test 'kglvqrro'", args{input: "kglvqrro", part: 1}, "DDUDRLRRUDRD"},
		{"test 'ulqzkmiv'", args{input: "ulqzkmiv", part: 1}, "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
		{"test 'ihgpwlah' p2", args{input: "ihgpwlah", part: 2}, "370"},
		{"test 'kglvqrro' p2", args{input: "kglvqrro", part: 2}, "492"},
		{"test 'ulqzkmiv' p2", args{input: "ulqzkmiv", part: 2}, "830"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d17MD5BFS(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d17MD5BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
