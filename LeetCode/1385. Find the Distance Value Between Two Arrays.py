class Solution:
    def findTheDistanceValue(self, arr1, arr2, d):
        count = 0
        for i in range(0, len(arr1)):
            flag = True
            for j in range(0, len(arr2)):
                if abs(arr1[i] - arr2[j]) <= d:
                    flag = False
                    break
            if flag:
                count += 1
        return count