package d14

import "testing"

func TestHash(t *testing.T) {
	type args struct {
		salt  string
		index int
		part  int
	}
	tests := []struct {
		name     string
		args     args
		wantHash string
	}{
		{"first triplet: 18", args{"abc", 18, 1}, "0034e0923cc38887a57bd7b1d4f953df"},
		{"next triplet: 39", args{"abc", 39, 1}, "347dac6ee8eeea4652c7476d0f97bee5"},
		{"next triplet: 92", args{"abc", 92, 1}, "ae2e85dd75d63e916a525df95e999ea0"},
		{"next triplet: 200", args{"abc", 200, 1}, "83501e9109999965af11270ef8d61a4f"},
		{"first triplet: 5", args{"abc", 5, 2}, "953a419067abb0b5e142680d73522236"},
		{"next triplet: 10", args{"abc", 10, 2}, "4a81e578d9f43511ab693eee1a75f194"},
		{"next triplet: 22859", args{"abc", 22859, 2}, "2e559978fffff9ac9c9012eb764c6391"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHash := Hash(tt.args.salt, tt.args.index, tt.args.part); gotHash != tt.wantHash {
				t.Errorf("Hash() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}
