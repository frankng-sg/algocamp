from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
        lines = f.readlines()
        f.close()
    return lines

def solve(lines: List[str]) -> int:
    time = [int(x) for x in lines[0][:-1].split(":")[1].strip().split(" ") if x != ""]
    dist = [int(x) for x in lines[1][:-1].split(":")[1].strip().split(" ") if x != ""]
    result = 1
    for i in range(len(time)):
        cnt = 0
        for t in range(1, time[i]):
            if t * (time[i] - t) > dist[i]:
                cnt += 1
        result *= cnt
    return result

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

