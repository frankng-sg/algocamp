class Solution:
    def minCostToMoveChips(self, chips) -> int:
        loc_dict = {}
        for loc in chips:
            if loc in loc_dict:
                loc_dict[loc] += 1
            else:
                loc_dict[loc] = 1

        min_cost = 100000
        for best_loc in loc_dict:
            cost = 0
            for loc in loc_dict:
                cost = cost + loc_dict[loc] * (abs(best_loc - loc) % 2)
            if cost < min_cost:
                min_cost = cost
        return min_cost


func = Solution()
chips = [2, 2, 2, 3, 3]  # Output: 2
chips = [6, 4, 7, 8, 2, 10, 2, 7, 9, 7]  # Output: 4
print(func.minCostToMoveChips(chips))
