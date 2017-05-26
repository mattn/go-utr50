package utr50

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "バッハはG線上のアリアを作曲した。",
			want:  "UTuUURUUUUUUUUUUUTu",
		},
		{
			input: "彼のミドルネームはSebastianだ。",
			want:  "UUUUUUTrUURRRRRRRRRUTu",
		},

		{
			input: "東京都港区芝公園四丁目二-八",
			want:  "UUUUUUUUUUUURU",
		},
		{
			input: "これはHewlett-Packardの製品だ。",
			want:  "UUURRRRRRRRRRRRRRRUUUUTu",
		},
	}
	for _, test := range tests {
		got := ""
		for _, r := range test.input {
			got += fmt.Sprint(Prop(r))
		}
		if got != test.want {
			t.Fatalf("want %q, got %q: %q", test.want, got, test.input)
		}
	}
}
