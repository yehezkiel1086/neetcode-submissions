func hasDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == tmp {
				return true
			}
		}
	}

	return false
}