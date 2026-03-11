package adventofcode

import (
	"reflect"
	"testing"
)

func Test_y16d21Move(t *testing.T) {
	type args struct {
		password []byte
		fields   []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "'bcdea' with 'move position 1 to position 4'", args: args{password: []byte("bcdea"), fields: []string{"move", "position", "1", "to", "position", "4"}}, want: []byte("bdeac")},
		{name: "'bdeac' with 'move position 3 to position 0'", args: args{password: []byte("bdeac"), fields: []string{"move", "position", "3", "to", "position", "0"}}, want: []byte("abdec")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d21Move(tt.args.password, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("y16d21Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d21Reverse(t *testing.T) {
	type args struct {
		password []byte
		fields   []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "'edcba' with 'reverse positions 0 through 4'", args: args{password: []byte("edcbafgh"), fields: []string{"reverse", "positions", "0", "through", "4"}}, want: []byte("abcdefgh")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d21Reverse(tt.args.password, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("y16d21Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d21Rotate(t *testing.T) {
	type args struct {
		password []byte
		fields   []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "'abcde' with 'rotate left 1 step'", args: args{password: []byte("abcde"), fields: []string{"rotate", "left", "1", "step"}}, want: []byte("bcdea")},
		{name: "'abdec' with 'rotate based on position of letter b'", args: args{password: []byte("abdec"), fields: []string{"rotate", "based", "on", "position", "of", "letter", "b"}}, want: []byte("ecabd")},
		{name: "'ecabd' with 'rotate based on position of letter d'", args: args{password: []byte("ecabd"), fields: []string{"rotate", "based", "on", "position", "of", "letter", "d"}}, want: []byte("decab")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d21Rotate(tt.args.password, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("y16d21Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d21Swap(t *testing.T) {
	type args struct {
		password []byte
		fields   []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "'abcde' with 'swap position 4 with position 0'", args: args{password: []byte("abcde"), fields: []string{"swap", "position", "4", "with", "position", "0"}}, want: []byte("ebcda")},
		{name: "'ebcda' with 'swap letter d with letter b'", args: args{password: []byte("ebcda"), fields: []string{"swap", "letter", "d", "with", "letter", "b"}}, want: []byte("edcba")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d21Swap(tt.args.password, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("y16d21Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d21(t *testing.T) {
	testInstructions := `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`
	type args struct {
		toScramble   string
		instructions string
		part         int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Part 1 - 'abcde' with test instructions", args: args{toScramble: "abcde", instructions: testInstructions, part: 1}, want: "decab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d21(tt.args.toScramble, tt.args.instructions, tt.args.part); got != tt.want {
				t.Errorf("y16d21() = %v, want %v", got, tt.want)
			}
		})
	}
}
