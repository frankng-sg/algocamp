class Solution:
    def bubbleSort(self, arrNum):
        for i in range(len(arrNum), 0, -1):
            for j in range(1, i):                
                if arrNum[j] < arrNum[j - 1]:
                    arrNum[j - 1], arrNum[j] = arrNum[j], arrNum[j - 1]
        return arrNum

    def findMinDiff(self, arrNum):
        min_diff = arrNum[-1]
        for i in range(1, len(arrNum)):
            if arrNum[i] - arrNum[i - 1] < min_diff:
                min_diff = arrNum[i] - arrNum[i - 1]
        return min_diff
                    
    def minimumAbsDifference(self, arrNum):
        self.bubbleSort(arrNum)
        min_diff = self.findMinDiff(arrNum)
        output = []
        for i in range(1, len(arrNum)):
            if arrNum[i] - arrNum[i - 1] == min_diff:
                output.append([arrNum[i - 1], arrNum[i]])
        return output
        
