package adventofcode

import (
	"reflect"
	"testing"
)

const y16d04TestInput string = `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

func Test_y16d04Decrypt(t *testing.T) {
	type args struct {
		name string
		id   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test decrypt", args{"qzmt-zixmtkozy-ivhz", 343}, "very encrypted name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d04Decrypt(tt.args.name, tt.args.id); got != tt.want {
				t.Errorf("y16d04Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d04RealRoom(t *testing.T) {
	type args struct {
		room y16d04Room
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test real room", args{y16d04Room{name: "aaaaa-bbb-z-y-x", checksum: "abxyz"}}, true},
		{"test real room", args{y16d04Room{name: "a-b-c-d-e-f-g-h", checksum: "abcde"}}, true},
		{"test real room", args{y16d04Room{name: "not-a-real-room", checksum: "oarel"}}, true},
		{"test real room", args{y16d04Room{name: "totally-real-room", checksum: "decoy"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d04RealRoom(tt.args.room); got != tt.want {
				t.Errorf("y16d04RealRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_y16d04ParseRoom(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantRoom y16d04Room
	}{
		{"test parse room", args{line: "aaaaa-bbb-z-y-x-343[abxyz]"}, y16d04Room{name: "aaaaa-bbb-z-y-x", checksum: "abxyz", id: 343}},
		{"test parse room", args{line: "a-b-c-d-e-f-g-h-987[abcde]"}, y16d04Room{name: "a-b-c-d-e-f-g-h", checksum: "abcde", id: 987}},
		{"test parse room", args{line: "not-a-real-room-404[oarel]"}, y16d04Room{name: "not-a-real-room", checksum: "oarel", id: 404}},
		{"test parse room", args{line: "totally-real-room-200[decoy]"}, y16d04Room{name: "totally-real-room", checksum: "decoy", id: 200}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoom := y16d04ParseRoom(tt.args.line); !reflect.DeepEqual(gotRoom, tt.wantRoom) {
				t.Errorf("y16d04ParseRoom() = %v, want %v", gotRoom, tt.wantRoom)
			}
		})
	}
}

func Test_y16d04(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test p1", args{y16d04TestInput, 1}, 1514},
		{"test p2", args{y16d04TestInput, 2}, 1514},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := y16d04(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("y16d04() = %v, want %v", got, tt.want)
			}
		})
	}
}
