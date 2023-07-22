function merge(left, right) {
    let arr = [];
    while (left.length && right.length) {
        if (left[0] < right[0]) arr.push(left.shift());
        else arr.push(right.shift());
    }
    return [...arr, ...left, ...right];
}

function mergeSort(arr) {
    if (arr.length <= 1) return arr;
    const half = arr.length >> 1;
    const left = arr.splice(0, half);
    return merge(mergeSort(left), mergeSort(arr));
}

console.log(mergeSort([3, 2, 1, 4, 4, 4, 4]));
