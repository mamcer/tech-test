package anagrama

import "sort"

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
