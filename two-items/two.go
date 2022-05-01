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
