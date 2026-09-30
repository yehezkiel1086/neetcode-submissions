func isValid(s string) bool {
	hash := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}
	stack := []rune{}

	for _, x := range s {
		// check if closing
		if op, ok := hash[x]; ok {
			// check if stack length is empty or is matching
			if len(stack) == 0 || stack[len(stack)-1] != op {
				return false
			}
			// remove opening from stack if closing
			stack = stack[:len(stack)-1]
		} else {
			// append to stack if opening
			stack = append(stack, x)
		}
	}

	return len(stack) == 0
}