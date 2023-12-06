map = dict({
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
    "1": "1",
    "2": "2",
    "3": "3",
    "4": "4",
    "5": "5",
    "6": "6",
    "7": "7",
    "8": "8",
    "9": "9",
    "0": "0"
})

def decode(s: str)->str:
    result = ""
    l, r = 0, 0
    n = len(s)
    while l < n:
        r += 1
        if s[l:r] in map:
            result += map[s[l:r]]
            l += 1
            r = 1
        elif r - l > 5 or r >= n:
            l += 1
            r = l
    return result

def solve()->int:
    with open('day-1-inp.txt') as f:
        lines = f.readlines()
        f.close()
    result = 0
    for line in lines:
        decoded = decode(line)
        # print(decoded)
        num = (ord(decoded[0]) - 48) * 10 + ord(decoded[-1]) - 48
        result += num
    return result
if __name__ == '__main__':
    print(solve())

# WRONG: 55725
