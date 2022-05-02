package palindrome

import "testing"

const palindromeExample = "racecar"
const notPalindromeExample = "mario"

var validateEquivalence = func(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestIsPalindromeFor(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		got := IsPalindromeFor(palindromeExample)
		validateEquivalence(t, got, true)
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsPalindromeFor(notPalindromeExample)
		validateEquivalence(t, got, false)
	})
}

func TestIsPalindromeSort(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		got := IsPalindromeSort(palindromeExample)
		validateEquivalence(t, got, true)
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsPalindromeSort(notPalindromeExample)
		validateEquivalence(t, got, false)
	})
}
