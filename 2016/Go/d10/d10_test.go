package d10

import "testing"

func TestPadNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{"test", args{"11"}, "00011"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := PadNumber(tt.args.s); gotOut != tt.wantOut {
				t.Errorf("PadNumber() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
