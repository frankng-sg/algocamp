from typing import List

def read_input()->List[str]:
    with open('day-2.txt') as f:
        lines = f.readlines()
        f.close()
    return lines

def solve(lines: List[str])->int:
    result = 0
    for line in lines:
        parts = line[:-1].split(':')
        game_id = int(parts[0].split(' ')[1])
        rounds = parts[1].split(';')
        max_cnt = dict({
            "red": 0,
            "green": 0,
            "blue": 0,
        })
        for round in rounds:
            for set in round.split(','):
                set = set.strip()
                cnt = int(set.split(' ')[0])
                color = set.split(' ')[1]
                if cnt > max_cnt[color]:
                    max_cnt[color] = cnt
        result += max_cnt["blue"] * max_cnt["red"] * max_cnt["green"]
    return result

if __name__ == '__main__':
    lines = read_input()
    print(solve(lines))
