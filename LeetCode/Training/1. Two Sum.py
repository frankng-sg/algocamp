###############################################################################
# ALGORITHM: Quick Sort without recursion
###############################################################################
# Input: an array "arr", a starting index "low", an ending index "high"
# Output: array "arr" with its elements from index low to index high sorted
# in ascending order
#
# Example:
#
# INPUT:
# arr = [3, 2, 9, 8, 5, 6, 0]
# low = 1
# high = 5
#
# OUTPUT:
# index  0  1  2  3  4  5  6
# arr = [3, 2, 5, 6, 8, 9, 0]
from random import randint

def partition(arr, index, low, high):
    # we find correct index for pivot element and split into two partitions:
    # -  left-partition contains elements whose value is smaller than or equal to
    #    pivot values
    # -  right-partition contains elements whose value is larger than
    #    pivot values

    # initially, we assume the pivot index be starting index that means
    # -  left-partition is empty
    # -  right-partition is at its max size
    pivot_index = low
    i = randint(low, high)
    index[i], index[high] = index[high], index[i]
    # we assume pivot value be the last element in the array
    # note: for different pivot value, please swap it with last element.
    pivot_value = arr[index[high]]
    # at this point, pivot value and pivot index are not synchronized
    # the loop below will find correct pivot index for pivot value
    for i in range(low, high):
        if arr[index[i]] <= pivot_value:
            # we move arr[i] to left-partition by swapping
            index[pivot_index], index[i] = index[i], index[pivot_index]
            # left-partition size grows by 1, so does pivot pivot_index
            pivot_index += 1

    # the correct pivot index has been found out which is next to the right of
    # finalized left-partition.
    # the element that sits at pivot_index is actually belong to right-partition
    # the element arr[high] is actually our pivot value
    # now we synchronize them.
    index[pivot_index], index[high] = index[high], index[pivot_index]

    return pivot_index

def quickSort(arr, index, low, high):
    # create a stack to mimic recursive strategy
    size = high - low + 1
    stack = [0] * size
    top = -1

    # push initial values
    top += 1
    stack[top] = low
    top += 1
    stack[top] = high

    # start quick sort algorithm
    while (top >= 0):
        # pop low and high values
        high = stack[top]
        top -= 1
        low = stack[top]
        top -= 1

        # set pivot element at correct position in the sub-array
        pivot = partition(arr, index, low, high)

        # now we have two partitions
        # left-partition contains elements from index low to pivot - 1
        # right-partition contains elements from index pivot + 1 to high

        # if left-partition is not empty, push it into stack
        if pivot - 1 > low:
            # push lower index in first
            top += 1
            stack[top] = low
            # so higher index is always higher
            top += 1
            stack[top] = pivot - 1

        # if left-partition is not empty, push it into stack
        if high > pivot + 1:
            # push lower index in first
            top += 1
            stack[top] = pivot + 1
            # so higher index is always higher
            top += 1
            stack[top] = high
# END ALGORITHM
###############################################################################

def binarySearch(target, arr, index, low, high):
    # create a stack to mimic recursive strategy
    stack = []

    # push initial values
    stack.append(low)
    stack.append(high)

    while stack != []: # while stack is not empty
        high = stack.pop()
        low = stack.pop()
        mid = int((low + high) / 2)

        if arr[index[mid]] == target:
            return mid

        # if target < mid node then we travel to left tree
        if target < arr[index[mid]] and mid > low:
            # push lower index in first
            stack.append(low)
            # so higher index is always higher
            stack.append(mid - 1)

        # if target > mid node then we travel to right tree
        elif target > arr[index[mid]] and high > mid:
            # push lower index in first
            stack.append(mid + 1)
            # so higher index is always higher
            stack.append(high)

    return -1

class Solution:
    def twoSum(self, arr, target):
        NOT_EXIST = -1
        # We first map position of all elements in array "arr" into
        # array "index".
        no_of_entries = len(arr)
        index = list(range(no_of_entries))
        # Then we sort array "arr" by manipulating its mapped index so that
        # tracing element position using array "index" gives us a sorted array
        quickSort(arr, index, 0, no_of_entries - 1)

        # Finally, for each element we will search if its partner exists
        # the sum of an element and its partner must be equal to target value
        for i in range(0, no_of_entries - 1):
            partner = binarySearch(target - arr[index[i]], arr, index, i + 1, no_of_entries - 1)
            if partner == NOT_EXIST:
                 continue
            else:
                return [index[i], index[partner]]
        return f"Cannot find two numbers whose sum is equal to {target}"
