class Solution:
    def largestRectangleArea(self, heights):
        """
        Test 1:
        Input: [2,1,5,6,2,3]
        Output: 10

        Test 2:
        Input: [2,3,4,5,2]
        Output: 10

        Test 3:
        Input: [3,2,4,5,2]
        Output: 10

        """

        heights = [-1] + heights + [-1]
        stack = [0]
        largest_area = 0
        for i in range(1, len(heights)):
            height_limit = heights[i]
            while heights[stack[-1]] > height_limit:
                area_height = heights[stack.pop()]
                area_starting_index = stack[-1]
                calc_area = (i - area_starting_index - 1) * area_height
                largest_area = max(largest_area, calc_area)

            stack.append(i)

        return largest_area


func = Solution()
heights = [2, 1, 5, 6, 2, 3]  # Output: 10
print(func.largestRectangleArea(heights))
heights = [2, 3, 4, 5, 2]  # Output: 10
print(func.largestRectangleArea(heights))
heights = [3, 2, 4, 5, 2]  # Output: 10
print(func.largestRectangleArea(heights))
