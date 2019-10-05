sequence = [0]

def binarySearch(low, high, key):
    # create a stack to mimic recursive strategy
    stack = []

    # push initial values
    stack.append(low)
    stack.append(high)

    while stack:  # while stack is not empty
        high = stack.pop()
        low = stack.pop()
        mid = int((low + high) / 2)

        if sequence[mid] == key:
            return mid
        if low + 1 == high:
            return high

        # if target < mid node then we travel to left tree
        if key < sequence[mid] and mid > low:
            # push lower index in first
            stack.append(low)
            # so higher index is always higher
            stack.append(mid - 1)

        # if target > mid node then we travel to right tree
        elif key > sequence[mid] and high > mid:
            # push lower index in first
            stack.append(mid + 1)
            # so higher index is always higher
            stack.append(high)

    # if algorithm does not return within the loop then key does not exist
    return -1


def generate(kth, size, num):
    global sequence
    if kth > size:
        return

    digit = num[len(num) - 1]
    if digit > '0':
        new_digit = str(ord(digit) - 1)
        new_num = num + new_digit
        sequence.append(kth + 1, size, int(new_num))
        generate(size, new_num)

    if digit < '9':
        new_digit = str(ord(digit) + 1)
        new_num = num + new_digit
        sequence.append(kth + 1, size, int(new_num))
        generate(size, new_num)


class Solution:
    def countSteppingNumbers(self, low: int, high: int) -> List[int]:
        global sequence
        for digit in range(1, 10):
            sequence.append(digit)
            generate(0, 9, str(digit))

        sequence = sorted(sequence)
        left = binarySearch(0, len(sequence) - 1, low)
        right = binarySearch(0, len(sequence) - 1, high)

        return sequence[left:right + 1]


func = Solution()
print("hello")
print(func.countSteppingNumbers(0, 21))
