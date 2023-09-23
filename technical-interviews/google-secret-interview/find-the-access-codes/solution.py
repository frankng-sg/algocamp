def bruteforce(nums):
    exist = dict()
    count = 0
    n = len(nums)
    for i in range(0, n - 2):
        for j in range(i + 1, n - 1):
            if nums[j] < nums[i] or nums[j] % nums[i] != 0:
                continue
            key2 = str(nums[i]) + " " + str(nums[j])
            if key2 in exist:
                continue
            exist[key2] = True
            for k in range(j + 1, n):
                if nums[k] < nums[j] or nums[k] % nums[j] != 0:
                    continue
                key3 = key2 + " " + str(nums[k])
                if key3 not in exist:
                    count += 1
                    exist[key3] = True
    return count


def solution(nums):
    return bruteforce(nums)


def stress_test():
    import random
    import time

    num_numbers = 3000
    seed = [1]
    for i in range(1, 31):
        seed.append(seed[i - 1] * 2)
    random_numbers = [
        seed[random.randint(0, len(seed) - 1)] for _ in range(num_numbers)
    ]
    start = time.time()
    solution(random_numbers)
    print("Stress Test Performance: ", time.time() - start, " seconds")


if __name__ == "__main__":
    print(solution([1, 2, 4, 8]))  # Output: 4
    print(solution([1, 2, 4, 1, 2, 4]))  # Output: 7
    print(solution([1, 2, 2, 2]))  # Output: 2
    print(solution([2, 2, 1, 2]))  # Output: 1
    print(solution([3, 4, 5]))  # Output: 0
    stress_test()
