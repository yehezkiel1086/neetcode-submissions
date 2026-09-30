func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tmp := map[rune]int{}

	for _, x := range s {
		tmp[x]++
	}

	for _, x := range t {
		tmp[x]--

		if tmp[x] < 0 {
			return false
		}
	}

	return true
}