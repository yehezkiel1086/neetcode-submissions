func twoSum(nums []int, target int) []int {
    hash := map[int]int{}

	for i, v := range nums {
		tmp := target - v
		if _, ok := hash[tmp]; ok {
			return []int{hash[tmp], i}
		}
		hash[v] = i
	}

	return []int{}
}
