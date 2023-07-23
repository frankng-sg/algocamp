class Set {
    constructor() {
        // Dictionary will hold the items of our set
        this.dictionary = {};
        this.length = 0;
    }

    // This method will check for the presence of an element and return true or false
    has(element) {
        return this.dictionary[element] !== undefined;
    }

    // This method will return all the values in the set
    values() {
        return Object.values(this.dictionary);
    }

    // Only change code below this line
    add(item) {
        if (this.has(item)) return false;
        this.dictionary[item] = item;
        return true;
    }
    remove(item) {
        if (!this.has(item)) return false;
        delete this.dictionary[item];
        return true;
    }
    size() {
        return Object.values(this.dictionary).length;
    }
    // Only change code above this line
}

const s = new Set();
console.log(s.add(1));
console.log(s.values());
console.log(s.add(1));
console.log(s.values());
