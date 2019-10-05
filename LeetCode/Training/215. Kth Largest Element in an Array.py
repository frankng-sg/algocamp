# Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order,
# not the kth distinct element.
#
# Example 1:
#
# Input: [3,2,1,5,6,4] and k = 2
# Output: 5
# Example 2:
#
# Input: [3,2,3,1,2,4,5,5,6] and k = 4
# Output: 4
# Note:
# You may assume k is always valid, 1 ≤ k ≤ array's length.

from random import randint


def partition(nums, left, right):
    """
    Algorithm: Partition by Tony Hoare
    """
    pivot = left + randint(0, right - left)
    nums[pivot], nums[right] = nums[right], nums[pivot]
    pivot = nums[right]

    store_index = left
    for i in range(left, right):
        if nums[i] > pivot:
            nums[i], nums[store_index] = nums[store_index], nums[i]
            store_index += 1
    nums[right], nums[store_index] = nums[store_index], nums[right]
    return store_index


def quick_select(nums, left, right, kth):
    """
    Algorithm: Quickselect by Tony Hoare
    """
    if left == right:  # list contains only one element
        return nums[left]  # then the element is kth largest
    pivot_index = partition(nums, left, right)
    if pivot_index is kth:
        return nums[pivot_index]
    if pivot_index < kth:
        return quick_select(nums, pivot_index + 1, right, kth)
    return quick_select(nums, left, pivot_index - 1, kth)


class Solution:

    def findKthLargest(self, nums, k):
        return quick_select(nums, 0, len(nums) - 1, k - 1)

# func = Solution()

# nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]
# k = 4
# expected_answer = 4
# print(f" Kth largest number = {func.findKthLargest(nums, k)}")

# nums = [3,2,1,5,6,4]
# k = 2
# expected_answer = 5
# print(f" Kth largest number = {func.findKthLargest(nums, k)}")
