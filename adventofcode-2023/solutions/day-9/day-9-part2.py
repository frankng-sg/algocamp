from math import lcm
from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
        lines = f.readlines()
        f.close()
    return lines

def zeroes(nums: List[int]) -> bool:
    for num in nums:
        if num != 0:
            return False
    return True

def solve(lines: List[str]) -> int:
    result = 0
    for line in lines:
        nums = [int(x) for x in line[:-1].split(" ")]
        sum_first_nums = nums[0]
        sign = -1
        while not zeroes(nums):
            next_nums = []
            for i in range(1, len(nums)):
                next_nums.append(nums[i] - nums[i-1])
            nums = next_nums
            sum_first_nums += sign * nums[0]
            sign = -sign
        result += sum_first_nums

    return result

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))
