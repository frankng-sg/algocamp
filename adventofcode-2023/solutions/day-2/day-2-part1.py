from typing import List

def read_input()->List[str]:
    with open('day-2.txt') as f:
        lines = f.readlines()
        f.close()
    return lines

def solve(lines: List[str])->int:
    max_cnt = dict({
        "red": 12,
        "green": 13,
        "blue": 14,
    })
    result = 0
    for line in lines:
        parts = line[:-1].split(':')
        game_id = int(parts[0].split(' ')[1])
        rounds = parts[1].split(';')
        possible = True
        for round in rounds:
            for set in round.split(','):
                set = set.strip()
                cnt = int(set.split(' ')[0])
                color = set.split(' ')[1]
                if cnt > max_cnt[color]:
                    possible = False
                    break
            if not possible:
                break
        if possible:
            result += game_id
    return result

if __name__ == '__main__':
    lines = read_input()
    print(solve(lines))
