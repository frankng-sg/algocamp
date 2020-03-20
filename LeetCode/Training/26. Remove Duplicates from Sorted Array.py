class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if nums == [] or nums is None:
            return 0

        count = 1
        i = 1
        while i < len(nums):
            if nums[i - 1] != nums[i]:
                count += 1
                i += 1
            else:
                nums.pop(i)
        return count


func = Solution()
print(func.removeDuplicates([1, 1, 2, 3, 4, 4]))  # Answer: 4
print(func.removeDuplicates([1]))  # Answer: 1
print(func.removeDuplicates([1, 1, 1, 1, 1, 1]))  # Answer: 1
