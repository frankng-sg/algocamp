function swap(arr, i, j) {
    let tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

function partition(arr, l, r) {
    let pivot = arr[r];
    let i = l - 1;
    let j;
    for (j = l; j <= r - 1; j++)
        if (arr[j] < pivot) {
            i += 1;
            swap(arr, i, j);
        }
    i += 1;
    swap(arr, i, r);
    return i;
}

function qSort(arr, l, r) {
    if (l >= r || l < 0) return;
    let p = partition(arr, l, r);
    qSort(arr, l, p - 1);
    qSort(arr, p + 1, r);
}

function quickSort(arr) {
    // Only change code below this line
    qSort(arr, 0, arr.length - 1);
    return arr;
    // Only change code above this line
}
console.log(quickSort([3, 2, 1, 4, 5, 4]));
