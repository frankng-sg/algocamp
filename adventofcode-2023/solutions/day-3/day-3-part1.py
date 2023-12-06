from typing import List


def read_input() -> List[str]:
    with open("day-3.txt") as f:
        lines = f.readlines()
        f.close()
    return lines


def is_part_number(map: List[str], irow: int, left: int, right: int) -> bool:
    def has_symbol_above(l: int, r: int) -> bool:
        if irow - 1 < 0:
            return False
        for i in range(l, r + 1):
            if map[irow - 1][i] != "." and not map[irow - 1][i].isdigit():
                return True
        return False

    def has_symbol_below(l: int, r: int) -> bool:
        if irow + 1 >= len(map):
            return False
        for i in range(l, r + 1):
            if map[irow + 1][i] != "." and not map[irow + 1][i].isdigit():
                return True
        return False

    def has_symbol_besides(l: int, r: int) -> bool:
        if map[irow][l] != "." and not map[irow][l].isdigit():
            return True
        if map[irow][r] != "." and not map[irow][r].isdigit():
            return True
        return False

    left -= 1
    if left < 0:
        left = 0
    if right >= len(map[irow]) - 1:
        right = len(map[irow]) - 2

    return (
        has_symbol_above(left, right)
        or has_symbol_below(left, right)
        or has_symbol_besides(left, right)
    )


def solve(lines: List[str]) -> int:
    result = 0
    irow = 0
    for line in lines:
        prev = -1
        for i in range(len(line) - 1):
            if not line[i].isdigit():
                if prev != -1:
                    num = int(line[prev:i])
                    if is_part_number(lines, irow, prev, i):
                        result += num
                    else:
                        print(num)
                    prev = -1
            elif prev == -1:
                prev = i
        if prev != -1:
            num = int(line[prev:])
            if is_part_number(lines, irow, prev, len(line)):
                result += num
            else:
                print(num)
        irow += 1
    return result


if __name__ == "__main__":
    lines = read_input()
    print(solve(lines))
