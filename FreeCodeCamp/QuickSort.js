function partition(arr, l, r) {
    if (l >= r) return;
    let pivot = arr[l + ((r - l) >> 1)];
    let i = l;
    let j = r;
    while (i <= j) {
        while (arr[i] < pivot) i++;
        while (arr[j] > pivot) j--;
        if (i <= j) {
            let tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
            i++;
            j--;
        }
    }
    partition(arr, l, j);
    partition(arr, i, r);
}

function quickSort(arr) {
    // Only change code below this line
    partition(arr, 0, arr.length - 1);
    return arr;
    // Only change code above this line
}

console.log(quickSort([3, 2, 4, 4, 5, 4]));
