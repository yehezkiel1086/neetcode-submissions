func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if !isAlphaNum(rune(s[i])) {
			i++
			continue
		}
		if !isAlphaNum(rune(s[j])) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsDigit(c) || unicode.IsLetter(c)
}