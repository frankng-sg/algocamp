def solve()->int:
    with open('day-1-inp.txt') as f:
        lines = f.readlines()
        f.close()

    sum = 0
    for line in lines:
        digit = -1
        for i in range(0, len(line)):
            if line[i] >= '0' and line[i] <= '9':
                digit = ord(line[i]) - 48
                break
        if digit == -1:
            continue
        num = digit
        for i in range(0, len(line)):
            if line[i] >= '0' and line[i] <= '9':
                digit = ord(line[i]) - 48
        num = num * 10 + digit
        sum += num
    return sum
if __name__ == '__main__':
    print(solve())
