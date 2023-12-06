from typing import List


def read_input() -> List[str]:
    with open("day-3.txt") as f:
        lines = f.readlines()
        f.close()
    return lines


def list_partno(map: List[str], irow: int, icol: int) -> List[int]:
    def extract(visited, i: int, j: int) -> int:
        left, right = j, j

        while left >= 0 and map[i][left].isdigit():
            left -= 1
        left += 1

        while right < len(map[i]) - 1 and map[i][right].isdigit():
            right += 1

        for k in range(left, right):
            visited.add(k)

        return int(map[i][left:right])

    def scan(i: int, l: int, r: int) -> List[int]:
        if i < 0 or i >= len(map):
            return []
        result = []
        visited = set()
        for j in range(l, r + 1):
            if j not in visited and map[i][j].isdigit():
                num = extract(visited, i, j)
                result.append(num)
        return result

    left, right = icol - 1, icol + 1
    if left < 0:
        left = 0
    if right >= len(map[irow]) - 1:
        right = len(map[irow]) - 2

    return (
        scan(irow - 1, left, right)
        + scan(irow, left, right)
        + scan(irow + 1, left, right)
    )


def solve(lines: List[str]) -> int:
    result = 0
    irow = 0
    for line in lines:
        for icol in range(len(line) - 1):
            if line[icol] == "*":
                nums = list_partno(lines, irow, icol)
                if len(nums) == 2:
                    result += nums[0] * nums[1]
        irow += 1
    return result


if __name__ == "__main__":
    lines = read_input()
    print(solve(lines))
