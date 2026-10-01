func plusOne(digits []int) []int {
    tt := 0

    for _, v := range digits {
        tt = tt*10 + v
    }

    tt++

    arr := []int{}
    for tt > 0 {
        arr = append([]int{tt % 10}, arr...)
        tt /= 10
    }

    return arr
}