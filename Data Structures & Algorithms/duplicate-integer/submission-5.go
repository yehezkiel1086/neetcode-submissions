func hasDuplicate(nums []int) bool {
    hash := map[int]int{}

	for _, v := range nums {
		if hash[v] >= 1 {
			return true
		}
		hash[v] += 1
	}

	return false
}
