class Solution:
    
    def uniqueOccurrences(self, arr):
        repeat = {}
        for num in arr:
            if (num in repeat):
                repeat[num] += 1
            else:
                repeat[num] = 1
        unique = set()
        for num in range(-1000, 1001):
            if num in repeat:
                if repeat[num] in unique:
                    return False
                else:
                    unique.add(repeat[num])
        return True
