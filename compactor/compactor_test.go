package compactor

import "testing"

var validateEquivalence = func(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestCompactFor(t *testing.T) {
	t.Run("compact should succeed", func(t *testing.T) {
		got := CompactFor("aaaaabbbadddccrrrrrrr")
		want := "a5b3a1d3c2r7"

		validateEquivalence(t, got, want)
	})

	t.Run("compact should succeed with blank", func(t *testing.T) {
		got := CompactFor("aaaaa    bbbadddcc")
		want := "a5 4b3a1d3c2"

		validateEquivalence(t, got, want)
	})

	t.Run("empty compact should succeed", func(t *testing.T) {
		got := CompactFor("")
		want := ""

		validateEquivalence(t, got, want)
	})

	t.Run("one empty compact should succeed", func(t *testing.T) {
		got := CompactFor("z")
		want := "z1"

		validateEquivalence(t, got, want)
	})
}
