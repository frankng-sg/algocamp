def solution(s):
    stack = 0
    count = 0
    for c in s:
        if c == ">":
            stack += 1
        elif c == "<":
            count += stack << 1
    return count


if __name__ == "__main__":
    print(solution(">----<"))  # Output: 2
    print(solution("<<>><"))  # Output: 4
    print(solution("><->--<"))  # Output: 6
    print(solution("-------"))  # Output: 0
