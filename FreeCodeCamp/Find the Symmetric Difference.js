function sym(...args) {
    const count = {};
    const symDiff = [];

    args.forEach((list) => {
        const uniqList = Array.from(new Set(list));
        uniqList.forEach((key) => {
            if (key in count) count[key] += 1;
            else count[key] = 1;
        });
    });

    for (const key in count) {
        if (count[key] % 2 == 1) {
            symDiff.push(parseInt(key));
        }
    }
    return symDiff;
}

console.log(sym([1, 2, 3], [5, 2, 1, 4])); // 3, 4, 5
console.log(sym([1, 2, 3, 3], [5, 2, 1, 4])); //3, 4, 5
console.log(sym([1, 2, 5], [2, 3, 5], [3, 4, 5])); //1, 4, 5
