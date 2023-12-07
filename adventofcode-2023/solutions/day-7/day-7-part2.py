import functools
from typing import List, Tuple

def read_input(filename) -> List[str]:
    with open(filename) as f:
        lines = f.readlines()
        f.close()
    return lines

def evaluate(hand: str) -> int:
    pair = []
    cnt = []
    for card in hand:
        if card != "J":
            found = False
            for i in range(len(pair)):
                if pair[i] == card:
                    cnt[i] += 1
                    found = True
            if not found:
                pair.append(card)
                cnt.append(1)

    cntJ = len(hand) - sum(cnt)
    if cntJ == len(hand):
        return 6 # five of a kind
    maxCnt = max(cnt)
    for i in range(len(cnt)):
        if cnt[i] == maxCnt:
            cnt[i] += cntJ
            break

    if len(pair) == 1:
        return 6 # five of a kind
    if len(pair) == 2:
        if cnt[0] == 4 or cnt[1] == 4:
            return 5 # four of a kind
        else:
            return 4 # full house
    if len(pair) == 3:
        if cnt[0] == 3 or cnt[1] == 3 or cnt[2] == 3:
            return 3 # three of a kind
        else:
            return 2 # two pair
    if len(pair) == 4:
        return 1 # one pair
    return 0 # high card

def solve(lines: List[str]) -> int:
    pwr = dict({
        "2": 0,
        "3": 1,
        "4": 2,
        "5": 3,
        "6": 4,
        "7": 5,
        "8": 6,
        "9": 7,
        "T": 8,
        "J": -1,
        "Q": 10,
        "K": 11,
        "A": 12
    })

    def card_cmp(a: str, b: str):
        for i in range(len(a)):
            if pwr[a[i]] > pwr[b[i]]:
                return +1
            if pwr[a[i]] < pwr[b[i]]:
                return -1
        return 0

    def hand_cmp(i: int, j: int):
        if strength[i] > strength[j]:
            return +1
        if strength[i] < strength[j]:
            return -1
        if card_cmp(hands[i], hands[j]) > 0:
            return +1
        if card_cmp(hands[i], hands[j]) < 0:
            return -1
        return 0

    hands = []
    strength = []
    bids = []
    idx = []
    for i in range(len(lines)):
        idx.append(i)
    for line in lines:
        [hand, bid] = line[:-1].split(" ")
        hands.append(hand)
        bids.append(int(bid))
        strength.append(evaluate(hand))

    idx.sort(key=functools.cmp_to_key(hand_cmp))

    result = 0
    for i in range(len(idx)):
        result += (i + 1) * bids[idx[i]]
    return result

if __name__ == "__main__":
    filename = ["sample.txt", "input.txt"]
    lines = read_input(filename[1])
    print(solve(lines))

