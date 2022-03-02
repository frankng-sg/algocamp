function countPoints(rings: string): number {
  const numberOfRods = 10;
  const numberOfRings = rings.length / 2;
  const rodWithRingColor = {};
  for (let i = 0; i < numberOfRings; i++) {
    const ringColor = rings[i * 2];
    const rodNumber = parseInt(rings[i * 2 + 1]);
    if (rodWithRingColor[rodNumber]) {
      rodWithRingColor[rodNumber][ringColor] = true;
    } else {
      rodWithRingColor[rodNumber] = { [ringColor]: true };
    }
  }
  let count = 0;
  for (let i = 0; i < numberOfRods; i++) {
    if (
      rodWithRingColor[i] &&
      rodWithRingColor[i]["R"] &&
      rodWithRingColor[i]["G"] &&
      rodWithRingColor[i]["B"]
    ) {
      count++;
    }
  }
  return count;
}

console.log(countPoints("B0B6G0R6R0R6G9"));
