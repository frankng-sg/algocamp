class Solution(object):
    def isPeak(self, nums, loc):
        if len(nums) == 1:
            return 0  # is peak

        if loc == 0:
            if nums[0] > nums[1]:
                return 0  # is peak
            return 1

        if loc == len(nums) - 1:
            if nums[loc] > nums[loc - 1]:
                return 0
            return -1

        if nums[loc - 1] < nums[loc] > nums[loc + 1]:
            return 0

        if nums[loc - 1] > nums[loc] > nums[loc + 1]:
            return -1  # is down-slope

        if nums[loc - 1] < nums[loc] < nums[loc + 1]:
            return 1  # is up-slope

        if nums[loc - 1] < nums[loc] > nums[loc + 1]:
            return -11  # is valley

    def findPeakElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        left, right = 0, len(nums)
        while left < right:
            mid = (left + right) // 2
            examine = self.isPeak(nums, mid)
            if examine == 0:
                return mid
            if examine == 1: # up-slope
                left = mid + 1
            else:
                right = mid

        return left
