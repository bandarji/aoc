package d05

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{"index 3231929", args{"abc", 3231929}, "1", true},
		{"index 5017308", args{"abc", 5017308}, "8", true},
		{"index 5278568", args{"abc", 5278568}, "f", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Hash(tt.args.s, tt.args.i)
			if got != tt.want {
				t.Errorf("Hash() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Hash() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBuildPassword(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantPassword string
	}{
		{"day 05 02", args{"abc"}, "05ace8e3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPassword := BuildPassword(tt.args.s); gotPassword != tt.wantPassword {
				t.Errorf("BuildPassword() = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}
