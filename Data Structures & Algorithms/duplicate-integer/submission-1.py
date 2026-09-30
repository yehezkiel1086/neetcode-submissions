class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash = {}

        for x in nums:
            hash[x] = hash.get(x, 0) + 1
            if hash[x] > 1:
                return True
        
        return False