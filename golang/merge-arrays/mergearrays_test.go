package mergearrays

import "testing"

var validateEquivalence = func(t *testing.T, got, want []int) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("len(got) != len(want)")
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("got:%v, want:%v", got, want)
			return
		}
	}
}

func TestIsPalindromeFor(t *testing.T) {
	t.Run("merge arrays", func(t *testing.T) {
		got := MergeArrays([]int{1, 3, 7}, []int{2, 5, 8, 9})
		validateEquivalence(t, got, []int{1, 2, 3, 5, 7, 8, 9})
	})
	t.Run("first on second", func(t *testing.T) {
		got := MergeArrays([]int{12, 14}, []int{2, 5, 8, 9})
		validateEquivalence(t, got, []int{2, 5, 8, 9, 12, 14})
	})
	t.Run("second on first", func(t *testing.T) {
		got := MergeArrays([]int{2, 5, 9}, []int{99, 110})
		validateEquivalence(t, got, []int{2, 5, 9, 99, 110})
	})
	t.Run("empty on non empty", func(t *testing.T) {
		got := MergeArrays([]int{}, []int{2, 5, 8, 9})
		validateEquivalence(t, got, []int{2, 5, 8, 9})
	})
	t.Run("non empty on empty", func(t *testing.T) {
		got := MergeArrays([]int{1, 3, 7}, []int{})
		validateEquivalence(t, got, []int{1, 3, 7})
	})
}
