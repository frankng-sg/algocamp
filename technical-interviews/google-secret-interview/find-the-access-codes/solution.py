def solution(l):
    count = 0
    size = len(l)
    if size < 3:
        return 0

    cache = [0] * size
    for x in range(size):
        for y in range(x + 1, size):
            if l[y] % l[x] == 0:
                cache[y] += 1
                count += cache[x]

    return count


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


_SEED = 1.23


def benchmark(sample_count):
    from random import seed, randint
    import timeit

    clock = timeit.default_timer

    seed(_SEED)
    samples = [
        [randint(1, 999999) for _ in range(randint(2, 2000))]
        for _ in range(sample_count)
    ]

    start = clock()
    for sample in samples:
        solution(sample)

    end = clock()
    print("%.4f s elapsed for %d samples." % (end - start, sample_count))


def test():
    # Provided test cases.
    assert solution([1, 1, 1]) == 1
    assert solution([1, 2, 3, 4, 5, 6]) == 3

    # Custom test cases.
    assert solution([1]) == 0
    assert solution([1, 2]) == 0
    assert solution([2, 4]) == 0
    assert solution([1, 1, 1, 1]) == 4
    assert solution([1, 1, 1, 1, 1]) == 10
    assert solution([1, 1, 1, 1, 1, 1]) == 20
    assert solution([1, 1, 1, 1, 1, 1, 1]) == 35
    assert solution([1, 1, 2]) == 1
    assert solution([1, 1, 2, 2]) == 4
    assert solution([1, 1, 2, 2, 2]) == 10
    assert solution([1, 1, 2, 2, 2, 3]) == 11
    assert solution([1, 2, 4, 8, 16]) == 10
    assert solution([2, 4, 5, 9, 12, 34, 45]) == 1
    assert solution([2, 2, 2, 2, 4, 4, 5, 6, 8, 8, 8]) == 90
    assert solution([2, 4, 8]) == 1
    assert solution([2, 4, 8, 16]) == 4
    assert solution([3, 4, 2, 7]) == 0
    assert solution([6, 5, 4, 3, 2, 1]) == 0
    assert solution([4, 7, 14]) == 0
    assert solution([4, 21, 7, 14, 8, 56, 56, 42]) == 9
    assert solution([4, 21, 7, 14, 56, 8, 56, 4, 42]) == 7
    assert solution([4, 7, 14, 8, 21, 56, 42]) == 4
    assert solution([4, 8, 4, 16]) == 2


def main():
    test()
    benchmark(100)


if __name__ == "__main__":
    main()
