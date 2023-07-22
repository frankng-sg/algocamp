function selectionSort(array) {
    // Only change code below this line
    for (let i = 0; i < array.length - 1; i++) {
        let minIndex = i;
        for (let j = i + 1; j < array.length; j++) {
            if (array[minIndex] > array[j]) minIndex = j;
        }
        if (minIndex != i) {
            let tmp = array[i];
            array[i] = array[minIndex];
            array[minIndex] = tmp;
        }
    }
    return array;
    // Only change code above this line
}

console.log(selectionSort([3, 2, 1, 4, 5, 4]));
