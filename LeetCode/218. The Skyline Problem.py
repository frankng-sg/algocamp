from typing import List


class Solution:
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        pass


func = Solution()
buildings = [[2, 9, 10], [3, 7, 15], [5, 12, 12], [15, 20, 10], [19, 24, 8]]
# Answer: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
print(func.getSkyline(buildings))

buildings = [[0, 2, 3], [2, 5, 3]]
# Answer: [[0,3],[5,0]]
print(func.getSkyline(buildings))

buildings = [[1, 2, 1], [1, 2, 2], [1, 2, 3]]
# Answer: [[1,3],[2,0]]
print(func.getSkyline(buildings))


def getSkyline_724ms_18MB(self, buildings: List[List[int]]) -> List[List[int]]:
    if buildings == []:
        return []

    LEFT = 0
    RIGHT = 1
    HEIGHT = 2

    rect = []
    for building in buildings:
        blendin = False
        for i in range(len(rect) - 1, -1, -1):
            if rect[i][RIGHT] < building[LEFT]:
                break

            if rect[i][LEFT] <= building[LEFT] < rect[i][RIGHT] < building[RIGHT]:
                blendin = True
                if rect[i][HEIGHT] == building[HEIGHT]:
                    rect[i][RIGHT] = building[RIGHT]
                    break
                elif rect[i][HEIGHT] > building[HEIGHT]:
                    building[LEFT] = rect[i][RIGHT]
                    rect.insert(i + 1, building)
                    break
                else:  # rect[i][HEIGHT] < building[HEIGHT]
                    rect[i][RIGHT] = building[LEFT]
                    rect.insert(i + 1, building)
            elif rect[i][LEFT] <= building[LEFT] < building[RIGHT] <= rect[i][RIGHT]:
                blendin = True
                if rect[i][HEIGHT] < building[HEIGHT]:
                    new_rect = [building[RIGHT], rect[i][RIGHT], rect[i][HEIGHT]]
                    rect[i][RIGHT] = building[LEFT]
                    rect.insert(i + 1, building)
                    rect.insert(i + 2, new_rect)
                    break
            elif building[LEFT] < rect[i][LEFT] < building[RIGHT] <= rect[i][RIGHT]:
                blendin = True
                if rect[i][HEIGHT] >= building[HEIGHT]:
                    building[RIGHT] = rect[i][LEFT]
                else:
                    rect[i][LEFT] = building[RIGHT]
            elif building[LEFT] <= rect[i][LEFT] < rect[i][RIGHT] <= building[RIGHT]:
                blendin = True
                if building[HEIGHT] == rect[i][HEIGHT]:
                    rect[i][RIGHT] = building[RIGHT]
                    building[RIGHT] = rect[i][LEFT]
                elif building[HEIGHT] > rect[i][HEIGHT]:
                    rect.pop(i)
                else:
                    new_rect = [rect[i][RIGHT], building[RIGHT], building[HEIGHT]]
                    rect.insert(i + 1, new_rect)
                    building[RIGHT] = rect[i][LEFT]
            elif rect[i][RIGHT] == building[LEFT] and rect[i][HEIGHT] == building[HEIGHT]:
                blendin = True
                rect[i][RIGHT] = building[RIGHT]
        if not blendin:
            rect.append(building)

    skyline = []
    size = len(rect)
    for i in range(size):
        if rect[i][LEFT] < rect[i][RIGHT]:
            skyline.append([rect[i][LEFT], rect[i][HEIGHT]])
            if i + 1 < size and rect[i][RIGHT] < rect[i + 1][LEFT]:
                skyline.append([rect[i][RIGHT], 0])
    skyline.append([rect[-1][RIGHT], 0])
    return skyline
