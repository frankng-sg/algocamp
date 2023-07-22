function pairwise(arr, arg) {
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

console.log(pairwise([1, 1, 1], 2));
console.log(pairwise([1, 4, 2, 3, 0, 5], 7));
