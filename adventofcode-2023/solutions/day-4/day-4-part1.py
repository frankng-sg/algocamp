from typing import List

def read_input() -> List[str]:
    with open("day-4.txt") as f:
        lines = f.readlines()
        f.close()
    return lines

def sanitize(line: str) -> str:
    return line.strip().replace("  ", " ")

def solve(lines: List[str]) -> int:
    result = 0
    for line in lines:
        parts = line[:-1].split(":")[1].split("|")
        winning = [int(x) for x in sanitize(parts[0]).split(" ")]
        onhand = [int(x) for x in sanitize(parts[1]).split(" ")]
        exist = dict()
        for x in onhand:
            exist[x] = True
        cnt = 0
        for x in winning:
            if x in exist:
                cnt += 1
        if cnt > 0:
            result += 2 ** (cnt - 1)
    return result

if __name__ == "__main__":
    lines = read_input()
    print(solve(lines))
