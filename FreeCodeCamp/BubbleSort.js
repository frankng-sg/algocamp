function bubbleSort(array) {
    // Only change code below this line
    let swapped = false;
    while (!swapped) {
        for (let i = 1; i < array.length; i++) {
            if (array[i - 1] > array[i]) {
                let tmp = array[i];
                array[i] = array[i - 1];
                array[i - 1] = tmp;
                swapped = true;
            }
        }
        if (!swapped) break;
        swapped = false;
    }
    return array;
    // Only change code above this line
}

console.log(bubbleSort([2, 1, 3, 5, 2]));
