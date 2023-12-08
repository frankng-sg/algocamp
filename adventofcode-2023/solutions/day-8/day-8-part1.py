import functools
from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
        lines = f.readlines()
        f.close()
    return lines

class Node:
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

def solve(lines: List[str]) -> int:
    nodes = dict()
    steps = lines[0][:-1]
    for i in range(2, len(lines)):
        line = lines[i][:-1].replace(" = (", " ").replace(", ", " ").replace(")", "")
        names = line.split(" ")
        for name in names:
            if name not in nodes:
                nodes[name] = Node(name)
        nodes[names[0]].left = nodes[names[1]]
        nodes[names[0]].right = nodes[names[2]]
    x = "AAA"
    target = "ZZZ"
    count = 0
    while True:
        for step in steps:
            if x == target:
                return count
            if step == "L":
                x = nodes[x].left.data
            elif step == "R":
                x = nodes[x].right.data
            count += 1

    return count

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

