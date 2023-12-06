from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
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
    k = 0
    while k < len(seeds):
        dest, big_range = seeds[k], seeds[k+1]
        for i in range(n):
            src = dest
            found = False
            for m in map[i]:
                if m[1] <= src and src < m[1] + m[2]:
                    small_range = min(m[1]+m[2]-src, big_range)
                    if small_range < big_range:
                        big_range = small_range
                    dest = m[0] + (src - m[1])
                    found = True
                    break
            if not found:
                dest = src
        if seeds[k+1] != big_range:
            seeds += [seeds[k] + big_range, seeds[k+1] - big_range]
            seeds[k+1] = big_range
        loc.append(dest)
        k += 2
    return min(loc)

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

