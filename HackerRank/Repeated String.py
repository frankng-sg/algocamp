def repeatedString(s, n):
    if s == "" or s is None or n == 0:
        return 0

    len_s = len(s)
    a_in_s = 0
    for letter in s:
        if letter == "a":
            a_in_s += 1
    result = a_in_s * (n // len_s)
    n = n % len_s
    for index in range(0, n):
        if s[index] == "a":
            result += 1
    return result


print(repeatedString("12aa32", 100))  # 34

print(repeatedString("abcac", 10))  # 4
print(repeatedString("bbcbc", 10))  # 0
print(repeatedString("aaaaa", 10))  # 10
