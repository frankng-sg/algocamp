from typing import List, Tuple

def read_input() -> List[str]:
    with open("day-5.txt") as f:
        lines = f.readlines()
        f.close()
    return lines

def read_seeds(lines: List[str])-> Tuple[List[int], List[str]]:
    data = lines[0].replace("seeds: ", "").replace("  ", " ").split(" ")
    return [int(x) for x in data], lines[2:]

def read_map(lines: List[str])-> Tuple[List[List[int]], List[str]]:
    i = 1
    result = []
    while i < len(lines) and lines[i] != "\n":
        result.append([int(x) for x in lines[i].strip().replace("  ", " ").split(" ")])
        i += 1
    return result, lines[i+1:]

def solve(lines: List[str]) -> int:
    seeds, lines = read_seeds(lines)
    n = 7
    map = []
    for _ in range(n):
        m, lines = read_map(lines)
        map.append(m)
    # print("seeds =", seeds)
    loc = []
    for seed in seeds:
        dest = seed
        for i in range(n):
            src = dest
            found = False
            for m in map[i]:
                if m[1] <= src and src <= m[1] + m[2]:
                    dest = m[0] + (src - m[1])
                    found = True
                    break
            if not found:
                dest = src
        loc.append(dest)
    return min(loc)

if __name__ == "__main__":
    lines = read_input()
    print(solve(lines))
