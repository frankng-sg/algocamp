/*
Given an array, return the first number that repeats

Example 1:
Given an array = [2, 5, 1, 2, 3, 5, 1, 2, 4]
It should return 2

Example 2:
Given an array = [2, 1, 1, 2, 3, 5, 1]
It should return 1

Example 3:
Given an array = [2, 3, 4, 5]
It should return undefined
 */
function getRepeatedNumber(arr) {
  const isRepeated = new Map();
  for (let i = 0; i < arr.length; i++) {
    if (isRepeated[arr[i]]) {
      return arr[i];
    }
    isRepeated[arr[i]] = true;
  }
}

console.log(getRepeatedNumber([2, 5, 1, 2, 3, 5, 1, 2, 4]));

console.log(getRepeatedNumber([2, 1, 1, 2, 3, 5, 1]));

console.log(getRepeatedNumber([2, 3, 4, 5]));
