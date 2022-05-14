package anagrama

import "testing"

var validateEquivalence = func(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestIsAnagramaSort(t *testing.T) {
	t.Run("is anagrama", func(t *testing.T) {
		got := IsAnagramaSort("mario", "omira")
		validateEquivalence(t, got, true)
	})

	t.Run("is not anagrama", func(t *testing.T) {
		got := IsAnagramaSort("mario", "orama")
		validateEquivalence(t, got, false)
	})

	t.Run("non empty vs empty should be false", func(t *testing.T) {
		got := IsAnagramaSort("mario", "")
		validateEquivalence(t, got, false)
	})

	t.Run("empty vs non empty should be false", func(t *testing.T) {
		got := IsAnagramaSort("", "mario")
		validateEquivalence(t, got, false)
	})

	t.Run("also even with spaces should be true", func(t *testing.T) {
		got := IsAnagramaSort("acuerdo japones", "ecuador esponja")
		validateEquivalence(t, got, true)
	})
}

func TestIsAnagramaMap(t *testing.T) {
	t.Run("is anagrama", func(t *testing.T) {
		got := IsAnagramaMap("mario", "omira")
		validateEquivalence(t, got, true)
	})

	t.Run("is not anagrama", func(t *testing.T) {
		got := IsAnagramaMap("mario", "orama")
		validateEquivalence(t, got, false)
	})

	t.Run("non empty vs empty should be false", func(t *testing.T) {
		got := IsAnagramaMap("mario", "")
		validateEquivalence(t, got, false)
	})

	t.Run("empty vs non empty should be false", func(t *testing.T) {
		got := IsAnagramaMap("", "mario")
		validateEquivalence(t, got, false)
	})

	t.Run("also even with spaces should be true", func(t *testing.T) {
		got := IsAnagramaMap("acuerdo japones", "ecuador esponja")
		validateEquivalence(t, got, true)
	})
}
