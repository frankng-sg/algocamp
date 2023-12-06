from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
        lines = f.readlines()
        f.close()
    return lines

def solve(lines: List[str]) -> int:
    time = int("".join([x for x in lines[0][:-1].split(":")[1].strip().split(" ")]))
    dist = int("".join([x for x in lines[1][:-1].split(":")[1].strip().split(" ")]))

    for left in range(1, time):
        if left * (time - left) > dist:
            break

    for right in range(time - 1, 0, -1):
        if right * (time - right) > dist:
            break

    return right - left + 1

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

