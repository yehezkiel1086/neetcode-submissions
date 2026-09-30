func isHappy(n int) bool {
    if n == 1 {
		return true
	}

	visit := map[int]bool{}

	for !visit[n] {
		visit[n] = true
		n = sum(n)
		if n == 1 {
			return true
		}
	}

	return false
}

func sum(n int) int {
	tt := 0

	for n > 0 {
		d := n%10
		d *= d
		tt += d
		n /= 10
	}

	return tt
}