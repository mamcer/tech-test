package compactor

import (
	"fmt"
	"strconv"
)

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

func ExtractSameChar(s string, si int) (string, int) {
	if si < len(s) {
		r := string(s[si])
		for i := si + 1; i < len(s); i++ {
			if rune(s[i]) != rune(r[0]) {
				return r, i
			} else {
				r += string(s[i])
			}
		}

		return r, len(s)
	}

	return " ", -1
}

func ExtractNumber(s string, si int) (int, int) {
	if si < len(s) {
		r := string(s[si])
		for i := si + 1; i < len(s); i++ {
			_, e := strconv.Atoi(string(s[i]))
			if e != nil {
				res, _ := strconv.Atoi(r)
				return res, i
			} else {
				r += string(s[i])
			}
		}

		res, _ := strconv.Atoi(r)
		return res, len(s)
	}

	return 0, -1
}

func UncompactFor(s string) string {
	if len(s) == 0 {
		return s
	}

	r := ""
	c := ""
	l := 0
	ni := 0
	for ni < len(s) && ni != -1 {
		c, ni = ExtractSameChar(s, ni)
		l, ni = ExtractNumber(s, ni)
		for j := 0; j < l; j++ {
			r += c
		}
	}

	return r
}
