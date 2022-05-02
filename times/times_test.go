package times

import "testing"

func TestStringTimes(t *testing.T) {
	t.Run("string exists", func(t *testing.T) {
		got := StringTimes("alabanza", "a")
		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("string does not exists", func(t *testing.T) {
		got := StringTimes("alabanza", "c")
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("empty string should return zero", func(t *testing.T) {
		got := StringTimes("", "a")
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
