def characterReplacement(s: str, k: int) -> int:
    count = {}
    l = 0
    maxf = 0
    for r in range(len(s)):
        count[s[r]] = 1 + count.get(s[r], 0)
        maxf = max(maxf, count[s[r]])
        if (r - l + 1) - maxf > k:
            count[s[l]] -= 1
            l += 1
    return r - l + 1


print(characterReplacement("ABAB", 2))  # Output: 4
print(characterReplacement("AABABBA", 1))  # Output: 4
print(characterReplacement("xAABBAABA", 2))  # Output: 6
print(characterReplacement("BAAAB", 2))  # Output: 5
print(characterReplacement("AAAxB", 2))  # Output: 5
print(characterReplacement("AAAAAAAAAabcdefghijklmnopAAAxB", 2))  # Output: 11
