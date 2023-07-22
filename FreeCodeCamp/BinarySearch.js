function binarySearch(searchList, value) {
    let arrayPath = [];
    let l = 0;
    let r = searchList.length - 1;
    while (l <= r) {
        let mid = l + ((r - l) >> 1);
        arrayPath.push(searchList[mid]);
        if (searchList[mid] == value) return arrayPath;
        if (searchList[mid] > value) r = mid - 1;
        else l = mid + 1;
    }
    return "Value Not Found";
}

console.log(binarySearch([0, 1, 2, 3, 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 49, 70], 1));
