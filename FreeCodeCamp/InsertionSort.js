function insertionSort(arr) {
    // Only change code below this line
    for (let i = 1; i < arr.length; i++) {
        let currVal = arr[i];
        let j;
        for (j = i - 1; j >= 0 && arr[j] > currVal; j--) arr[j + 1] = arr[j];
        arr[j + 1] = currVal;
    }
    return arr;
    // Only change code above this line
}

console.log(insertionSort([2, 1, 3, 5, 2]));
