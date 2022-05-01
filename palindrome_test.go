package palindrome

import "testing"

func TestIsPalindromeFor(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		got := IsPalindromeFor("racecar")
		want := true
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsPalindromeFor("mario")
		want := false
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})
}

func TestIsPalindromeSort(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		got := IsPalindromeSort("racecar")
		want := true
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsPalindromeSort("mario")
		want := false
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})
}
