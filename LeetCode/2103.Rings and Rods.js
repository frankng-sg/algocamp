function countPoints(rings) {
    var _a;
    var numberOfRods = 10;
    var numberOfRings = rings.length / 2;
    var rodWithRingColor = {};
    for (var i = 0; i < numberOfRings; i++) {
        var ringColor = rings[i * 2];
        var rodNumber = parseInt(rings[i * 2 + 1]);
        if (rodWithRingColor[rodNumber]) {
            rodWithRingColor[rodNumber][ringColor] = true;
        }
        else {
            rodWithRingColor[rodNumber] = (_a = {}, _a[ringColor] = true, _a);
        }
    }
    var count = 0;
    for (var i = 0; i < numberOfRods; i++) {
        if (rodWithRingColor[i] &&
            rodWithRingColor[i]["R"] &&
            rodWithRingColor[i]["G"] &&
            rodWithRingColor[i]["B"]) {
            count++;
        }
    }
    return count;
}
console.log(countPoints("B0B6G0R6R0R6G9"));
