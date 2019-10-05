def binarySearch(target, arr, low, high):
    # create a stack to mimic recursive strategy
    stack = []

    # push initial values
    stack.append(low)
    stack.append(high)

    while stack:  # while stack is not empty
        high = stack.pop()
        low = stack.pop()
        mid = int((low + high) / 2)

        if arr[mid] == target:
            return mid

        # if target < mid node then we travel to left tree
        if target < arr[mid] and mid > low:
            # push lower index in first
            stack.append(low)
            # so higher index is always higher
            stack.append(mid - 1)

        # if target > mid node then we travel to right tree
        elif target > arr[mid] and high > mid:
            # push lower index in first
            stack.append(mid + 1)
            # so higher index is always higher
            stack.append(high)

    return -1

class Solution:
    def twoSum(self, arr, target):
        NOT_EXIST = -1

        no_of_entries = len(arr)

        # Finally, for each element we will search if its partner exists
        # the sum of an element and its partner must be equal to target value
        for i in range(0, no_of_entries - 1):
            partner = binarySearch(target - arr[i], arr, i + 1, no_of_entries - 1)
            if partner == NOT_EXIST:
                 continue
            else:
                return [i+1, partner+1]
        return f"Cannot find two numbers whose sum is equal to {target}"
