func twoSum(nums []int, target int) []int {
    hash := map[int]int{}

	for i, v := range nums {
		diff := target - v
		if j, ok := hash[diff]; ok {
			return []int{j, i}
		}
		hash[v] = i
	}

	return []int{0, 0}
}
