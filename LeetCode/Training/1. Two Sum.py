class Solution:
    def twoSum(self, list_of_nums, target):
        no_of_entries = len(list_of_nums)

        for firstEntry in range(0, no_of_entries - 1):
            for secondEntry in range(firstEntry + 1, no_of_entries):
                if list_of_nums[firstEntry] + list_of_nums[secondEntry] == target:
                    return (firstEntry, secondEntry)

        return f"Cannot find two numbers whose sum is equal to {target}" 