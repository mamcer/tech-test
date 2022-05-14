package anagrama

import "testing"

var validateEquivalence = func(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestIsAnagrama(t *testing.T) {
	got := IsAnagrama("mario", "oiram")

	validateEquivalence(t, got, true)
}
