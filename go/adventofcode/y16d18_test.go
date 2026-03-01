package adventofcode

import (
	"reflect"
	"testing"
)

func Test_y16d18MakeRow(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantRow []bool
	}{
		{"test '..^^.'", args{input: "..^^."}, []bool{false, false, true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRow := y16d18MakeRow(tt.args.input); !reflect.DeepEqual(gotRow, tt.wantRow) {
				t.Errorf("y16d18MakeRow() = %v, want %v", gotRow, tt.wantRow)
			}
		})
	}
}

func Test_y16d18NextRow(t *testing.T) {
	type args struct {
		row []bool
	}
	tests := []struct {
		name        string
		args        args
		wantNextRow []bool
	}{
		{"test '..^^.'", args{row: y16d18MakeRow("..^^.")}, []bool{false, true, true, true, true}},
		{"test '.^^^^'", args{row: y16d18MakeRow(".^^^^")}, []bool{true, true, false, false, true}},
		{"test '^^^^..^^^.'", args{row: y16d18MakeRow("^^^^..^^^.")}, []bool{true, false, false, true, true, true, true, false, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNextRow := y16d18NextRow(tt.args.row); !reflect.DeepEqual(gotNextRow, tt.wantNextRow) {
				t.Errorf("y16d18NextRow() = %v, want %v", gotNextRow, tt.wantNextRow)
			}
		})
	}
}

func Test_y16d18CountSafeTiles(t *testing.T) {
	type args struct {
		row []bool
	}
	tests := []struct {
		name          string
		args          args
		wantSafeTiles int
	}{
		{"test '..^^.'", args{row: y16d18MakeRow("..^^.")}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSafeTiles := y16d18CountSafeTiles(tt.args.row); gotSafeTiles != tt.wantSafeTiles {
				t.Errorf("y16d18CountSafeTiles() = %v, want %v", gotSafeTiles, tt.wantSafeTiles)
			}
		})
	}
}

func Test_y16d18(t *testing.T) {
	type args struct {
		input    string
		rowCount int
	}
	tests := []struct {
		name          string
		args          args
		wantSafeTiles int
	}{
		{"test '..^^.'", args{input: "..^^.", rowCount: 3}, 6},
		{"test '.^^.^.^^^^'", args{input: ".^^.^.^^^^", rowCount: 10}, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSafeTiles := y16d18(tt.args.input, tt.args.rowCount); gotSafeTiles != tt.wantSafeTiles {
				t.Errorf("y16d18() = %v, want %v", gotSafeTiles, tt.wantSafeTiles)
			}
		})
	}
}
