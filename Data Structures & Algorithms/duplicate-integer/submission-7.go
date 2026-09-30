func hasDuplicate(nums []int) bool {
    hash := map[int]int{}

	for _, v := range nums {
		if hash[v] != 0 {
			return true
		}
		hash[v]++
	}

	return false
}
