package sns_go_namehash

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		sns     string
		wantOk  bool
		wantSNS string
	}{
		{"1234.seedao", true, "1234.seedao"},
		{"1234ab.seedao", true, "1234ab.seedao"},
		{"abcd.seedao", true, "abcd.seedao"},
		{"AbcD.seedao", true, "abcd.seedao"},
		{"!abc", true, "!abc"},
		{"$abc", true, "$abc"},
		{"(abc", true, "(abc"},
		{")abc", true, ")abc"},
		{"*abc", true, "*abc"},
		{"+abc", true, "+abc"},
		{"-abc", true, "-abc"},
		{"_abc", true, "_abc"},
		//
		{"#abc", false, ""},
		{"%abc", false, ""},
		{"'abc", false, ""},
		{":abc", false, ""},
		{"@abc", false, ""},
		{"[abc", false, ""},
		{"<abc", false, ""},
		{"{abc", false, ""},
		{"|abc", false, ""},
	}

	for _, tt := range tests {
		ok, sns := Normalize(tt.sns)
		if ok != tt.wantOk {
			t.Errorf("Normalize(%s)'s ok got: %v, want:%v", tt.sns, ok, tt.wantOk)
		}
		if tt.wantOk && sns != tt.wantSNS {
			t.Errorf("Normalize(%s)'s sns got:%s, want:%s", tt.sns, sns, tt.wantSNS)
		}
	}
}

func TestNamehash(t *testing.T) {
	tests := []struct {
		sns  string
		want string
	}{
		{"seedao",
			"0x5e55419d79fa352b3401db837903c9d6425f83393880fd079b57ad5f232def51"},
		{"baiyu.seedao",
			"0xf6f4c6b561092c29a877e94011088977faca0eaddc5654d0f036599de86bcdc4"},
	}

	for _, tt := range tests {
		got := Namehash(tt.sns)
		if got != tt.want {
			t.Errorf("Namehash(%s) = %s, but want %s", tt.sns, got, tt.want)
		}
	}
}
