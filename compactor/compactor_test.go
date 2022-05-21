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

func TestExtractSegment(t *testing.T) {
	t.Run("Extract should succeed", func(t *testing.T) {
		got1, got2 := ExtractSameChar("aaaaabbbcc", 0)
		if got1 != "aaaaa" {
			t.Errorf("got:%v, want:%v\n", got1, "aaaaa")
		}
		if got2 != 5 {
			t.Errorf("got:%v, want:%v\n", got1, 5)
		}

		got1, got2 = ExtractSameChar("aaaaabbbcc", 5)
		if got1 != "bbb" {
			t.Errorf("got:%v, want:%v\n", got1, "bbb")
		}
		if got2 != 8 {
			t.Errorf("got:%v, want:%v\n", got1, 5)
		}

		got1, got2 = ExtractSameChar("aaaaabbbcc", 8)
		if got1 != "cc" {
			t.Errorf("got:%v, want:%v\n", got1, "cc")
		}
		if got2 != 10 {
			t.Errorf("got:%v, want:%v\n", got2, 10)
		}

		got1, got2 = ExtractSameChar("a10", 0)
		if got1 != "a" {
			t.Errorf("got:%v, want:%v\n", got1, "a")
		}
		if got2 != 1 {
			t.Errorf("got:%v, want:%v\n", got2, 1)
		}

		got3, got4 := ExtractNumber("a10", 1)
		if got3 != 10 {
			t.Errorf("got:%v, want:%v\n", got1, 10)
		}
		if got4 != 3 {
			t.Errorf("got:%v, want:%v\n", got2, 3)
		}
	})
}

func TestUncompactFor(t *testing.T) {
	t.Run("uncompact should succeed", func(t *testing.T) {
		got := UncompactFor("a5b3a1d3c2r7")
		want := "aaaaabbbadddccrrrrrrr"

		validateEquivalence(t, got, want)
	})

	t.Run("uncompact one char should succeed", func(t *testing.T) {
		got := UncompactFor("a1")
		want := "a"

		validateEquivalence(t, got, want)
	})

	t.Run("uncompact one char multiple times should succeed", func(t *testing.T) {
		got := UncompactFor("a10")
		want := "aaaaaaaaaa"

		validateEquivalence(t, got, want)
	})

	t.Run("uncompact empty should succeed", func(t *testing.T) {
		got := UncompactFor("")
		want := ""

		validateEquivalence(t, got, want)
	})
}
