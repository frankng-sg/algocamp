# 5079. Intersection of Three Sorted Arrays
# Given three integer arrays arr1, arr2 and arr3 sorted in strictly increasing order,
# return a sorted array of only the integers that appeared in all three arrays.
# Constraints:
# 1 <= arr1.length, arr2.length, arr3.length <= 1000
# 1 <= arr1[i], arr2[i], arr3[i] <= 2000
class Solution:
    def arraysIntersection(self, arr1, arr2, arr3):
        set2 = set(arr2)
        set3 = set(arr3)
        output = []
        for num in arr1:
            if num in arr2 and num in arr3:
                output.append(num)
        return output

# func = Solution()
# arr1 = [1,2,3,4,5]
# arr2 = [1,2,5,7,9]
# arr3 = [1,3,4,5,8]
# print(func.arraysIntersection(arr1, arr2, arr3))
