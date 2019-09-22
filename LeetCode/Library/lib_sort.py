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

def partition(arr, low, high):
    # we find correct index for pivot element and split into two partitions:
    # -  left-partition contains elements whose value is smaller than or equal to
    #    pivot values
    # -  right-partition contains elements whose value is larger than
    #    pivot values

    # initially, we assume the pivot index be starting index that means
    # -  left-partition is empty
    # -  right-partition is at its max size
    pivot_index = low
    # we use random pivot element for optimal performance and then swap it
    # to the last element. The reason is our partition algorithm only works
    # with last element as the pivot value
    i = randint(low, high)
    arr[i], arr[high] = arr[high], arr[i]

    pivot_value = arr[high]
    # at this point, pivot value and pivot index are not synchronized
    # the loop below will find correct pivot index for pivot value
    for i in range(low, high):
        if arr[i] <= pivot_value:
            # we move arr[i] to left-partition by swapping
            arr[pivot_index], arr[i] = arr[i], arr[pivot_index]
            # left-partition size grows by 1, so does pivot pivot_index
            pivot_index += 1

    # the correct pivot index has been found out which is next to the right of
    # finalized left-partition.
    # the element that sits at pivot_index is actually belong to right-partition
    # the element arr[high] is actually our pivot value
    # now we synchronize them.
    arr[pivot_index], arr[high] = arr[high], arr[pivot_index]

    return pivot_index

def quickSort(arr, low, high):
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
        pivot = partition(arr, low, high)

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

# Test ALGORITHM
nums = [8, 7, 5, 4, 3, 1, 2, 6, 9]
quickSort(nums, 0, len(nums) - 1)
print(nums)
