###############################################################################
# ALGORITHM: Binary Search without recursion
###############################################################################
# Input: an array "arr"
#        a starting index "low"
#        an ending index "high"
#        a value to be searched "key"
# Output: return -1 if key does not exist
#         return index of the key found in the array
#
# Example:
#
# INPUT:
# arr = [3, 2, 9, 8, 5, 6, 0]
# low = 1
# high = 5
# key = 8
#
# OUTPUT:
# 3 (because arr[3] == 8)

def binarySearch(arr, low, high, key):
    # create a stack to mimic recursive strategy
    stack = []

    # push initial values
    stack.append(low)
    stack.append(high)

    while stack != []: # while stack is not empty
        high = stack.pop()
        low = stack.pop()
        mid = int((low + high) / 2)

        if arr[mid] == key:
            return mid

        # if target < mid node then we travel to left tree
        if key < arr[mid] and mid > low:
            # push lower index in first
            stack.append(low)
            # so higher index is always higher
            stack.append(mid - 1)

        # if target > mid node then we travel to right tree
        elif key > arr[mid] and high > mid:
            # push lower index in first
            stack.append(mid + 1)
            # so higher index is always higher
            stack.append(high)

    # if algorithm does not return within the loop then key does not exist
    return -1
# END ALGORITHM
###############################################################################

arr = [1, 2, 3, 4, 5, 6]
print(binarySearch(arr, 0, 5, 1))
print(binarySearch(arr, 0, 5, 10))
