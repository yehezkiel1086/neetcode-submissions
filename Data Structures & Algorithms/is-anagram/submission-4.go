func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash_s := map[rune]int{}

	for _, v := range s {
		hash_s[v]++
	}

	for _, v := range t {
		hash_s[v]--
		if hash_s[v] < 0 {
			return false
		}
	}

	return true
}
