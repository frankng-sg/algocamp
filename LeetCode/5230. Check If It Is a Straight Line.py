class Solution:
    def checkStraightLine(self, xylist) -> bool:
        try:
            origin = xylist[0]
            slope = abs(xylist[1][1] - origin[1]) / abs(xylist[1][0] - origin[0])
            for i in range(2, len(xylist)):
                new_slope = abs(xylist[i][1] - origin[1]) / abs(xylist[i][0] - origin[0])
                if slope != new_slope:
                    return False
            return True
        except:
            return False

func = Solution()
xylist = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
print(func.checkStraightLine(xylist))