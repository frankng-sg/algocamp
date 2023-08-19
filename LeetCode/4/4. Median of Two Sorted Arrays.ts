function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    const merged = [];
    let i = 0;
    let j = 0;
    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] < nums2[j]) {
            merged.push(nums1[i]);
            i++;
        } else {
            merged.push(nums2[j]);
            j++;
        }
    }
    if (i < nums1.length) {
        merged.push(...nums1.slice(i));
    }
    if (j < nums2.length) {
        merged.push(...nums2.slice(j));
    }

    const mid = Math.floor(merged.length / 2);
    if (merged.length % 2 == 1) {
        return merged[mid];
    }
    return (merged[mid] + merged[mid - 1]) / 2;
}