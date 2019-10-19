def sort(sequence):
    list_char = list(sequence)
    length = len(list_char)

    for i in range(length, 0, -1):
        for j in range(1, i):
            if list_char[j - 1] > list_char[j]:
                temp = list_char[j]
                list_char[j] = list_char[j - 1]
                list_char[j - 1] = temp

    output = ''.join(list_char)

    return output


def main():
    sequence = input()
    print(sort(sequence))


if __name__ == '__main__':
    main()
