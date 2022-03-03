function kthDistinct(arr: string[], k: number): string {
  const recurrence = {};
  for (let i = 0; i < arr.length; i++) {
    if (recurrence[arr[i]] === undefined) {
      recurrence[arr[i]] = 1;
    } else {
      recurrence[arr[i]] += 1;
    }
  }
  let count = 0;
  for (let i = 0; i < arr.length; i++) {
    if (recurrence[arr[i]] == 1) {
      count++;
      if (count == k) {
        return arr[i];
      }
    }
  }
  return "";
}
