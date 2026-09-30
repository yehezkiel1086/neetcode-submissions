class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash = {}

        for x in nums:
            if x in hash:
                return True
            hash[x] = hash.get(x, 0) + 1
        
        return False