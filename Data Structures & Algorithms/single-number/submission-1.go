func singleNumber(nums []int) int {
	seen := map[int]bool{}

	for _, v := range nums {
		if seen[v] {
			delete(seen, v)
		} else {
			seen[v] = true
		}
	}

	for num := range seen {
		return num
	}

	return -1
}
