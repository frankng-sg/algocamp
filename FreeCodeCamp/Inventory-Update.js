function updateInventory(curInv, newInv) {
    let curInvMap = new Map();

    curInv.forEach((item) => {
        curInvMap.set(item[1], item[0]);
    });
    newInv.forEach((item) => {
        if (curInvMap.has(item[1])) curInvMap.set(item[1], curInvMap.get(item[1]) + item[0]);
        else curInvMap.set(item[1], item[0]);
    });
    let updatedInv = Array.from(curInvMap, ([key, val]) => [val, key]);
    updatedInv.sort((a, b) => (a[1] >= b[1] ? 1 : -1));
    return updatedInv;
}

// Example inventory lists
var curInv = [
    [21, "Bowling Ball"],
    [2, "Dirty Sock"],
    [1, "Hair Pin"],
    [5, "Microphone"]
];

var newInv = [
    [2, "Hair Pin"],
    [3, "Half-Eaten Apple"],
    [67, "Bowling Ball"],
    [7, "Toothpaste"]
];

console.log(updateInventory(curInv, newInv));
