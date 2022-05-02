package two

import "testing"

var validateEquivalence = func(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func TestTwoItemsNested(t *testing.T) {

	t.Run("two items success", func(t *testing.T) {
		got := TwoItemsNested([]int{6, 1, 12, 64, 4, 2, 13, 10}, 18)

		validateEquivalence(t, got, true)
	})

	t.Run("two items fail", func(t *testing.T) {
		got := TwoItemsNested([]int{6, 1, 12, 64, 4, 2, 13, 10}, 100)
		validateEquivalence(t, got, false)
	})

	t.Run("empty array fail", func(t *testing.T) {
		got := TwoItemsNested([]int{}, 10)
		validateEquivalence(t, got, false)
	})
}

func TestTwoItemsSort(t *testing.T) {

	t.Run("two items success", func(t *testing.T) {
		got := TwoItemsSort([]int{7, 1, 3, 2, 6}, 5)

		validateEquivalence(t, got, true)
	})

	t.Run("two items fail", func(t *testing.T) {
		got := TwoItemsSort([]int{6, 1, 12, 64, 4, 2, 13, 10}, 100)
		validateEquivalence(t, got, false)
	})

	t.Run("empty array fail", func(t *testing.T) {
		got := TwoItemsSort([]int{}, 10)
		validateEquivalence(t, got, false)
	})
}

func TestTwoItemsMap(t *testing.T) {

	t.Run("two items success", func(t *testing.T) {
		got := TwoItemsMap([]int{7, 1, 3, 2, 6}, 5)

		validateEquivalence(t, got, true)
	})

	t.Run("two items fail", func(t *testing.T) {

		got := TwoItemsMap([]int{6, 1, 12, 64, 4, 2, 13, 10}, 100)
		validateEquivalence(t, got, false)
	})

	t.Run("empty array fail", func(t *testing.T) {
		got := TwoItemsMap([]int{}, 10)
		validateEquivalence(t, got, false)
	})
}
