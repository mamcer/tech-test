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
