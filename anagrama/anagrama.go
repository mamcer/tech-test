package anagrama

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsAnagramaSort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)

	return s1 == s2
}

func IsAnagramaMap(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var m1 map[string]int = make(map[string]int)
	var m2 map[string]int = make(map[string]int)

	for i := 0; i < len(s1); i++ {
		if _, ok := m1[string(s1[i])]; ok {
			m1[string(s1[i])] += 1
		} else {
			m1[string(s1[i])] = 1
		}

		if _, ok := m2[string(s2[i])]; ok {
			m2[string(s2[i])] += 1
		} else {
			m2[string(s2[i])] = 1
		}
	}

	for i := range m1 {
		if _, ok := m2[i]; ok {
			if m1[i] != m2[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
