func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l+r)/2
		if target == nums[m] {
			return m
		} else if target < nums[m] {
			r = m-1
		} else {
			l = m+1
		}
	}

	return -1
}
