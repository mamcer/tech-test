package times

func StringTimes(s string, v string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == v {
			c += 1
		}
	}

	return c
}

func StringTimesMap(s string, v string) (map[string]int, int) {
	var m map[string]int = make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			m[string(s[i])] += 1
		} else {
			m[string(s[i])] = 1
		}
	}

	return m, m[v]
}
