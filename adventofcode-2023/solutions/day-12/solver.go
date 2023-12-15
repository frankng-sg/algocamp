package main

func countCombo4(record []RecordType, reqs []int) int {
    cache := make(map[[2]int]int)
    var try func(int, int) int
    try = func(loc, req int) int {
        key := [2]int{loc, req}
        if cnt, ok := cache[key]; ok {
            return cnt
        }
        cache[key] = 0
        for i := loc; i < len(record); i++ {
            if canBePlaced(record, i, i+reqs[req]-1) {
                if req == len(reqs)-1 {
                    if noMoreDamaged(record, i+reqs[req]+1) {
                        cache[key] += 1
                    }
                } else {
                    cache[key] += try(i+reqs[req]+1, req+1)
                }
            }
            if record[i] == DAMAGED {
                break
            }
        }
        return cache[key]
    }
    return try(0, 0)
}

func countCombo3(record []RecordType, reqs []int) int {
    count := 0
    limit := make([]int, len(reqs)+1)
    for i := len(reqs) - 1; i >= 0; i-- {
        limit[i] = limit[i+1] + 1 + reqs[i]
    }
    var try func(int, int)
    try = func(k, from int) {
        if k >= len(reqs) {
            if noMoreDamaged(record, from) {
                count += 1
            }
            return
        }
        for i := from; i < len(record)-limit[k+1]; i++ {
            if canBePlaced(record, i, i+reqs[k]-1) {
                try(k+1, i+reqs[k]+1)
            }
            if record[i] == DAMAGED {
                break
            }
        }
    }
    try(0, 0)
    return count
}

func countCombo2(record []RecordType, reqs []int) int {
    var try func(int, int, int)
    result := 0
    limit := make([]int, len(reqs)+1)
    for i := len(reqs) - 1; i >= 0; i-- {
        limit[i] = limit[i+1] + 1 + reqs[i]
    }
    try = func(i, cnt, idx int) {
        if i > 0 {
            if record[i-1] == DAMAGED {
                if cnt == 0 {
                    idx++
                }
                if idx >= len(reqs) {
                    return
                }
                cnt++
                if cnt > reqs[idx] {
                    return
                }
            } else {
                if cnt > 0 && cnt != reqs[idx] {
                    return
                }
                cnt = 0
            }
            if i > len(record)-limit[idx+1]+1 {
                return
            }
        }
        if i >= len(record) {
            if ((record[i-1] == DAMAGED && cnt == reqs[idx]) || record[i-1] == OPERATIONAL) &&
                idx == len(reqs)-1 {
                result += 1
            }
            return
        }
        if record[i] != UNKNOWN {
            try(i+1, cnt, idx)
            return
        }
        for _, v := range []RecordType{OPERATIONAL, DAMAGED} {
            record[i] = v
            try(i+1, cnt, idx)
        }
        record[i] = UNKNOWN
    }
    try(0, 0, -1)
    return result
}

func countCombo(record []RecordType, reqs []int) int {
    var try func(int)
    count := 0
    try = func(i int) {
        if i >= len(record) {
            if isPossibleCombo(record, reqs) {
                count += 1
            }
            return
        }

        if record[i] != UNKNOWN {
            try(i + 1)
            return
        }
        for _, v := range []RecordType{OPERATIONAL, DAMAGED} {
            record[i] = v
            try(i + 1)
        }
        record[i] = '?'
    }
    try(0)
    return count
}
