func hasDuplicate(nums []int) bool {
    hash := make(map[int]int)

    for _, v := range nums {
        if _, ok := hash[v]; ok {
            return true
        }
        hash[v] = 1
    }

    return false
}
