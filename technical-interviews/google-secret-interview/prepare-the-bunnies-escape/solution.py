def solution(map):
    h = len(map)
    w = len(map[0])
    min_path = w * h

    def bfs(x, y, min_path):
        path = []
        for _ in range(h):
            path.append([0] * w)
        queue = [[x, y]]
        path[x][y] = 1
        while len(queue) > 0:
            [x, y] = queue.pop(0)

            if path[x][y] >= min_path:
                return min_path

            if x == h - 1 and y == w - 1:
                return path[x][y]

            if x + 1 < h and path[x + 1][y] == 0 and map[x + 1][y] == 0:
                path[x + 1][y] = path[x][y] + 1
                queue.append([x + 1, y])

            if x - 1 >= 0 and path[x - 1][y] == 0 and map[x - 1][y] == 0:
                path[x - 1][y] = path[x][y] + 1
                queue.append([x - 1, y])

            if y + 1 < w and path[x][y + 1] == 0 and map[x][y + 1] == 0:
                path[x][y + 1] = path[x][y] + 1
                queue.append([x, y + 1])

            if y - 1 >= 0 and path[x][y - 1] == 0 and map[x][y - 1] == 0:
                path[x][y - 1] = path[x][y] + 1
                queue.append([x, y - 1])

    for i in range(h):
        for j in range(w):
            if map[i][j] == 1:
                map[i][j] = 0
                min_path = bfs(0, 0, min_path)
                map[i][j] = 1
    min_path = bfs(0, 0, min_path)
    return min_path


if __name__ == "__main__":
    assert (
        solution(
            [
                [0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1],
                [0, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0],
            ]
        )
        == 11
    )
