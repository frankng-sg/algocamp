from typing import List

def read_input() -> List[str]:
    with open("day-4.txt") as f:
        lines = f.readlines()
        f.close()
    return lines

def sanitize(line: str) -> str:
    return line.strip().replace("  ", " ")

def solve(lines: List[str]) -> int:
    points = []
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
        points.append(cnt)
    n = len(lines)
    cards = [0] * n
    for i in range(n):
        cards[i] += 1
        if points[i] > 0:
            for j in range(i + 1, min(n, i + points[i] + 1)):
                cards[j] += cards[i]
    print(cards)
    return sum(cards)

if __name__ == "__main__":
    lines = read_input()
    print(solve(lines))
