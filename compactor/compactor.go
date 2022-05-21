package compactor

import "fmt"

func CompactFor(s string) string {
	if len(s) == 0 {
		return s
	}

	if len(s) == 1 {
		return s + "1"
	}

	r := ""
	p := rune(s[0])
	c := 1
	for i := 1; i < len(s); i++ {
		if rune(s[i]) != rune(p) {
			r += string(p) + fmt.Sprintf("%v", c)
			p = rune(s[i])
			c = 1
		} else {
			c += 1
		}
	}
	r += string(p) + fmt.Sprintf("%v", c)

	return r
}
