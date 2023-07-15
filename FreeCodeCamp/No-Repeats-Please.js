function permAlone(str) {
    function hasRepeatingChars(chars) {
        for (let i = chars.length - 1; i > 0; i--) if (chars[i] == chars[i - 1]) return true;
        return false;
    }

    function genAll(chars, n) {
        if (n <= 0) {
            if (!hasRepeatingChars(chars)) count++;
            return;
        }
        for (let i = n; i >= 0; i--) {
            [chars[i], chars[n]] = [chars[n], chars[i]];
            genAll(chars, n - 1);
            [chars[i], chars[n]] = [chars[n], chars[i]];
        }
    }

    let count = 0;
    const chars = str.split("");
    genAll(chars, str.length - 1);
    return count;
}

console.log(permAlone("abcdefa"));
