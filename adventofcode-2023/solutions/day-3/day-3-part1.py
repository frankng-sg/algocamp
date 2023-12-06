from typing import List

def read_input()->List[str]:
    with open('./day-3.txt') as f:
        lines = f.readlines()
        f.close()
    return lines

def is_part_number(map: List[str], irow: int, left: int, right: int)->bool:
    scan_left, scan_right = left - 1, right
    if scan_left < 0:
        scan_left = 0
    if scan_right >= len(map[irow]):
        scan_right = len(map[irow]) - 1
    if scan_left < left and map[irow][scan_left] != '.' :
        return True
    if scan_right >= right and  map[irow][scan_right] != '.':
        return True
    if irow - 1 >= 0:
        for i in range(scan_left, scan_right + 1):
            if not map[irow - 1][i].isdigit() and map[irow - 1][i] != '.':
                return True
    if irow + 1 < len(map):
        for i in range(scan_left, scan_right + 1):
            if not map[irow + 1][i].isdigit() and map[irow + 1][i] != '.':
                return True
    return False

def solve(lines: List[str])->int:
    result = 0
    irow = 0
    for line in lines:
        line = line[:-1]
        prev = -1
        for i in range(len(line)):
            if prev == -1 and line[i].isdigit() and (i == 0 or not line[i - 1].isdigit()):
                    prev = i
            elif not line[i].isdigit() and line[i - 1].isdigit():
                if is_part_number(lines, irow, prev, i):
                    num = int(line[prev:i])
                    result += num
                else:
                    print(line[prev:i])
                prev = -1
        if prev != -1:
            if is_part_number(lines, irow, prev, len(line)):
                num = int(line[prev:])
                result += num
            else:
                print(line[prev:i])
        irow += 1
    return result

if __name__ == '__main__':
    lines = read_input()
    print(solve(lines))

# Ans: 524857 - TOO HIGH
