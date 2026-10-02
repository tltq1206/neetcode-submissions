func validPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	for l < r {
		if s[l] != s[r]{
			if isPalindrome(s, l+1, r) || isPalindrome(s, l, r - 1) {
				return true
			}
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNumericRegex(b byte) bool {
	return (b >= 'a' && b <='z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <='9')
}

func isPalindrome( s string ,l, r int) bool {
	for l < r {
		if !isAlphaNumericRegex(s[l]) {
			l++
			continue
		}
		if !isAlphaNumericRegex(s[r]){
			r--
			continue
		}
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}