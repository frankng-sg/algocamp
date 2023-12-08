from math import lcm
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

def reached(srcs: List[str]) -> bool:
    for src in srcs:
        if src[-1] != "Z":
            return False
    return True

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
    srcs = []
    for name in nodes.keys():
        if name[-1] == "A":
            srcs.append(name)
    req_steps = []
    for i in range(len(srcs)):
        count = 0
        j = 0
        while srcs[i][-1] != "Z":
            if steps[j] == "L":
                srcs[i] = nodes[srcs[i]].left.data
            elif steps[j] == "R":
                srcs[i] = nodes[srcs[i]].right.data
            count += 1
            j += 1
            if j == len(steps):
                j = 0
        req_steps.append(count)
    return lcm(*req_steps) # use python3.9

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

