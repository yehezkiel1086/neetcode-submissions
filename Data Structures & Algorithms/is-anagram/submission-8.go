func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[rune]int{}

	for _, v := range s {
		hash[v]++
	}

	for _, v := range t {
		if hash[v] <= 0 {
			return false
		}
		hash[v]--
	}

	return true
}
