class Solution:
    def isValid(self, s: str) -> bool:
        hash = {
            ')': '(',
            '}': '{',
            ']': '[',
        }

        stack = []

        for x in s:
            if x in hash:
                if len(stack) == 0 or stack.pop() != hash[x]:
                    return False
            elif x in hash.values():
                stack.append(x)

        return len(stack) == 0