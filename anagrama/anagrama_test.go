package anagrama

import "testing"

var validateEquivalence = func(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestIsAnagramaSort(t *testing.T) {

	t.Run("is palindrome", func(t *testing.T) {
		got := IsAnagramaSort("mario", "omira")
		validateEquivalence(t, got, true)
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsAnagramaSort("mario", "orama")
		validateEquivalence(t, got, false)
	})

	t.Run("empty is false", func(t *testing.T) {
		got := IsAnagramaSort("mario", "")
		validateEquivalence(t, got, false)
	})

	t.Run("empty is false", func(t *testing.T) {
		got := IsAnagramaSort("", "mario")
		validateEquivalence(t, got, false)
	})

	t.Run("is palindrome", func(t *testing.T) {
		got := IsAnagramaSort("acuerdo japones", "ecuador esponja")
		validateEquivalence(t, got, true)
	})
}
