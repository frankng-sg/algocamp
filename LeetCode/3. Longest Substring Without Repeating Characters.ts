function isDuplicated(cnt: object): boolean {
    return Object.keys(cnt).some((key) => cnt[key] > 1);
}

function lengthOfLongestSubstring(s: string): number {
    let [out_len, out_l, out_r] = [0, 0, 0];
    let [left, right] = [0, 0];
    const lim_r = s.length;
    const cnt = {};
    while (right <= lim_r) {
        if (isDuplicated(cnt)) {
            cnt[s[left]]--;
            left++;
            if (left >= lim_r - out_len) break;
        } else {
            if (right - left > out_len) {
                out_len = right - left;
                [out_l, out_r] = [left, right];
            }
            if (s[right] in cnt) {
                cnt[s[right]]++;
            } else {
                cnt[s[right]] = 1;
            }
            right++;
        }
    }

    return out_len;
}
