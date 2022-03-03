function destCity(paths) {
    var isNotDestination = {};
    for (var i = 0; i < paths.length; i++) {
        isNotDestination[paths[i][0]] = true;
    }
    for (var i = 0; i < paths.length; i++) {
        if (isNotDestination[paths[i][0]] == undefined) {
            return paths[i][0];
        }
        if (isNotDestination[paths[i][1]] == undefined) {
            return paths[i][1];
        }
    }
    return "";
}
// Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
// Output: "Sao Paulo"
console.log(destCity([
    ["London", "New York"],
    ["New York", "Lima"],
    ["Lima", "Sao Paulo"],
]));
// Input: paths = [["B","C"],["D","B"],["C","A"]]
// Output: "A"
console.log(destCity([
    ["B", "C"],
    ["D", "B"],
    ["C", "A"],
]));
// Input: paths = [["A","Z"]]
// Output: "Z"
console.log(destCity([["A", "Z"]]));
