package compactor

import "testing"

func TestCompactFor(t *testing.T) {
	got := CompactFor("aaaaabbbadddcc")
	want := "a5b3a1d3c2"
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}
