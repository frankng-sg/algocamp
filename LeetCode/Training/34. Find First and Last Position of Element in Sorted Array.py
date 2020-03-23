class Solution:
    def searchRange(self, nums, target):
        not_found = [-1, -1]

        if not nums:
            return not_found

        low, high = 0, len(nums) - 1
        while low < high:
            mid = (low + high) // 2
            if nums[mid] >= target:
                high = mid
            else:
                low = mid + 1
        if nums[low] != target:
            return not_found
        left = low

        high = len(nums) - 1
        while low < high - 1:
            mid = (low + high) // 2
            if target < nums[mid]:
                high = mid - 1
            else:
                low = mid

        right = high if nums[high] == target else low

        return [left, right]

ans  = Solution()
print(ans.searchRange([5,7,7,8,8,10], 8))
print(ans.searchRange([2, 2], 2))