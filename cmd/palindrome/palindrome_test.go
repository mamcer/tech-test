package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		got := IsPalindrome("racecar")
		want := true
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})

	t.Run("is not palindrome", func(t *testing.T) {
		got := IsPalindrome("mario")
		want := false
		if got != want {
			t.Errorf("Got:%v, expect:%v", got, want)
		}
	})
}
