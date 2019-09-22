class Solution:
    def partition(self, arr, l, h): 
        i = ( l - 1 ) 
        x = arr[h] 
      
        for j in range(l, h): 
            if   arr[j] <= x: 
      
                # increment index of smaller element 
                i = i + 1
                arr[i], arr[j] = arr[j], arr[i] 
      
        arr[i + 1], arr[h] = arr[h], arr[i + 1] 
        return (i + 1) 
      
    # Function to do Quick sort 
    # arr[] --> Array to be sorted, 
    # l  --> Starting index, 
    # h  --> Ending index 
    def quickSort(self, arr, l, h): 
      
        # Create an auxiliary stack 
        size = h - l + 1
        stack = [0] * (size) 
      
        # initialize top of stack 
        top = -1
      
        # push initial values of l and h to stack 
        top = top + 1
        stack[top] = l 
        top = top + 1
        stack[top] = h 
      
        # Keep popping from stack while is not empty 
        while top >= 0: 
      
            # Pop h and l 
            h = stack[top] 
            top = top - 1
            l = stack[top] 
            top = top - 1
      
            # Set pivot element at its correct position in 
            # sorted array 
            p = self.partition( arr, l, h ) 
      
            # If there are elements on left side of pivot, 
            # then push left side to stack 
            if p-1 > l: 
                top = top + 1
                stack[top] = l 
                top = top + 1
                stack[top] = p - 1
      
            # If there are elements on right side of pivot, 
            # then push right side to stack 
            if p + 1 < h: 
                top = top + 1
                stack[top] = p + 1
                top = top + 1
                stack[top] = h 
      

    def findMinDiff(self, arrNum):
        min_diff = arrNum[-1]
        for i in range(1, len(arrNum)):
            if arrNum[i] - arrNum[i - 1] < min_diff:
                min_diff = arrNum[i] - arrNum[i - 1]
        return min_diff
                    
    def minimumAbsDifference(self, arrNum):
        self.quickSort(arrNum, 0, len(arrNum) - 1)
        min_diff = self.findMinDiff(arrNum)
        output = []
        for i in range(1, len(arrNum)):
            if arrNum[i] - arrNum[i - 1] == min_diff:
                output.append([arrNum[i - 1], arrNum[i]])
        return output
        
