package palindrome

func IsPalindromeFor(s string) bool {
	if s == "" {
		return false
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func IsPalindromeForHalf(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func invertString(s string) string {
	var r string = ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return r
}

func IsPalindromeSort(s string) bool {
	if s == "" {
		return false
	}

	si := invertString(s)
	return s == si
}
