def subtract(x, y, k, base):
    out = ""
    carry = 0
    for i in range(k - 1, -1, -1):
        digit_x = ord(x[i]) - 48
        digit_y = ord(y[i]) - 48 + carry
        carry = 0
        if digit_x < digit_y:
            carry = 1
            digit_x += base
        out = chr(digit_x - digit_y + 48) + out
    return out


def solution(n, b):
    k = len(n)
    exist = dict()
    loop = 0
    while n not in exist:
        loop += 1
        exist[n] = loop
        y = "".join(sorted(n))  # asceding order
        x = y[::-1]  # descending order
        n = subtract(x, y, k, b)
    return loop - exist[n] + 1


if __name__ == "__main__":
    print(solution("1211", 10))  # Output: 1
    print(solution("210022", 3))  # Output: 3
