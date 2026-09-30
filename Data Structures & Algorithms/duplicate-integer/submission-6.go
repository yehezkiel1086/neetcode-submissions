func hasDuplicate(nums []int) bool {
    hash := map[int]int{}

	for _, v := range nums {
		hash[v]++
		if hash[v] > 1 {
			return true
		}
	}

	return false
}
