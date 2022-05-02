package times

import "testing"

var validateEquivalence = func(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestStringTimes(t *testing.T) {
	t.Run("string exists", func(t *testing.T) {
		got := StringTimes("alabanza", "a")

		validateEquivalence(t, got, 4)
	})

	t.Run("string does not exists", func(t *testing.T) {
		got := StringTimes("alabanza", "c")

		validateEquivalence(t, got, 0)
	})

	t.Run("empty string should return zero", func(t *testing.T) {
		got := StringTimes("", "a")

		validateEquivalence(t, got, 0)
	})
}
