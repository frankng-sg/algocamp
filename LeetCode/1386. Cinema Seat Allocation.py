class Solution:
    def maxNumberOfFamilies(self, n, reservedSeats):
        reservedSeats.sort(key=lambda x: x[0])
        reservedSeats.append([0, 0])
        M = len(reservedSeats)
        free = [True] * 11
        free[reservedSeats[0][1]] = False
        i = 1
        count = 0
        while i < M:
            if reservedSeats[i][0] == reservedSeats[i - 1][0]:
                free[reservedSeats[i][1]] = False
            else:
                mid_seats = True
                if free[2] and free[3] and free[4] and free[5]:
                    mid_seats = False
                    count += 1
                if free[6] and free[7] and free[8] and free[9]:
                    mid_seats = False
                    count += 1
                if mid_seats and free[4] and free[5] and free[6] and free[7]:
                    count += 1
                n -= 1
                free = [True] * 11
                free[reservedSeats[i][1]] = False
            i += 1

        return count + n * 2


ans = Solution()
print(ans.maxNumberOfFamilies(3, [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]))