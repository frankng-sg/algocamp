function pairwise1(arr, arg) {
    let sum = 0;
    let used = {};
    for (let i = 0; i < arr.length; i++) {
        if (!(i in used)) {
            for (let j = i + 1; j < arr.length; j++) {
                if (!(j in used) && arr[i] + arr[j] == arg) {
                    sum += i + j;
                    used[j] = true;
                    break;
                }
            }
        }
    }
    return sum;
}

function pairwise2(arr, arg) {
    let sum = 0;
    let pair = {};
    let count = {};
    let used = {};
    for (let i = 0; i < arr.length; i++) {
        let v = arg - arr[i];
        if (!([v, 0] in pair)) {
            pair[[v, 0]] = i;
            count[v] = 1;
        } else {
            pair[[v, count[v]]] = i;
            count[v] += 1;
        }
        used[i] = false;
    }
    for (let i = 0; i < arr.length; i++) {
        if (!used[i]) {
            let v = arr[i];
            for (let j = 0; j < count[v]; j++) {
                let pairIndex = pair[[v, j]];
                if (pairIndex != i && !used[pairIndex]) {
                    sum += i + pairIndex;
                    used[pairIndex] = true;
                    used[i] = true;
                    break;
                }
            }
        }
    }
    return sum;
}

console.log(pairwise1([1, 1, 1], 2));
console.log(pairwise1([1, 4, 2, 3, 0, 5], 7));

console.log(pairwise2([1, 1, 1], 2));
console.log(pairwise2([1, 4, 2, 3, 0, 5], 7));
