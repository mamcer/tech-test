package two

import (
	"sort"
)

func TwoItemsNested(items []int, value int) bool {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i]+items[j] == value {
				return true
			}
		}
	}

	return false
}

func TwoItemsSort(items []int, value int) bool {
	sort.Ints(items)

	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i]+items[j] == value {
				return true
			}

			if items[i]+items[j] > value {
				break
			}
		}
	}

	return false
}

func TwoItemsMap(items []int, value int) bool {
	// sort.Ints(items)
	var here map[int]int = make(map[int]int)

	for i := 0; i < len(items); i++ {
		// if items[i] > value {
		// 	break
		// }
		me := items[i]

		if v, ok := here[me]; ok {
			if v == 0 {
				return true
			}
		}

		here[me] = 1

		want := value - me
		if v, ok := here[want]; ok {
			if v == 1 {
				return true
			}
		}

		here[want] = 0
	}

	return false
}
