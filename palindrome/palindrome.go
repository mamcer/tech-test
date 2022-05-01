package palindrome

func IsPalindromeFor(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
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
	si := invertString(s)
	return s == si
}
