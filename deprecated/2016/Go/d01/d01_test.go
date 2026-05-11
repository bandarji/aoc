package d01

import (
	"testing"
)

func TestNormalizePosition(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0,0", args{0, 0}, "0,0"},
		{"-1,0", args{-1, 0}, "-1,0"},
		{"0,-1", args{0, -1}, "0,-1"},
		{"-1,-1", args{-1, -1}, "-1,-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePosition(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("NormalizePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInstruction(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name      string
		args      args
		wantTurn  string
		wantSteps int
	}{
		{"L50", args{"L50"}, "L", 50},
		{"R50", args{"R50"}, "R", 50},
		{"L5", args{"L5"}, "L", 5},
		{"R5", args{"R5"}, "R", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTurn, gotSteps := ParseInstruction(tt.args.instruction)
			if gotTurn != tt.wantTurn {
				t.Errorf("ParseInstruction() gotTurn = %v, want %v", gotTurn, tt.wantTurn)
			}
			if gotSteps != tt.wantSteps {
				t.Errorf("ParseInstruction() gotSteps = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}

func TestTurn(t *testing.T) {
	type args struct {
		turn   string
		facing int
	}
	tests := []struct {
		name          string
		args          args
		wantDirection int
	}{
		{"L, 0", args{"L", 0}, 3},
		{"L, 1", args{"L", 1}, 0},
		{"L, 2", args{"L", 2}, 1},
		{"L, 3", args{"L", 3}, 2},
		{"R, 0", args{"R", 0}, 1},
		{"R, 1", args{"R", 1}, 2},
		{"R, 2", args{"R", 2}, 3},
		{"R, 3", args{"R", 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDirection := Turn(tt.args.turn, tt.args.facing); gotDirection != tt.wantDirection {
				t.Errorf("Turn() = %v, want %v", gotDirection, tt.wantDirection)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		sx int
		sy int
		ex int
		ey int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"distance 4/1", args{0, 0, 2, 2}, 4},
		{"distance 4/2", args{-1, -1, 1, 1}, 4},
		{"distance 10/1", args{0, 0, 6, 4}, 10},
		{"distance 10/2", args{-5, 0, 0, 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.sx, tt.args.sy, tt.args.ex, tt.args.ey); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		v1 int
		v2 int
	}
	tests := []struct {
		name  string
		args  args
		wantA int
	}{
		{"positive/positive", args{1, 11}, 10},
		{"positive/negative", args{1, -9}, 10},
		{"negative/negative", args{-4, -14}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := Abs(tt.args.v1, tt.args.v2); gotA != tt.wantA {
				t.Errorf("Abs() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestDay(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 blocks", args{"R2, L3", 1}, 5},
		{"2 blocks", args{"R2, R2, R2", 1}, 2},
		{"12 blocks", args{"R5, L5, R5, R3", 1}, 12},
		{"4 blocks", args{"R8, R4, R4, R8", 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}
