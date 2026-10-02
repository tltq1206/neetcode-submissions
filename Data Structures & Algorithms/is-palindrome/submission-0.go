func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlphanumericRegex(s[l]) {
			l++
			continue
		}
		if !isAlphanumericRegex(s[r]) {
			r--
			continue
		}
		if !strings.EqualFold(string(s[l]), string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumericRegex(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
