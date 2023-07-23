var called = 0;
var hash = (string) => {
    called++;
    var hashed = 0;
    for (var i = 0; i < string.length; i++) {
        hashed += string.charCodeAt(i);
    }
    return hashed;
};
var HashTable = function () {
    this.collection = {};
    // Only change code below this line
    this.add = function (key, val) {
        const hashed = hash(key);

        if (!(hashed in this.collection)) this.collection[hashed] = [];
        this.collection[hashed].push([key, val]);
    };
    this.lookup = function (key) {
        const hashed = hash(key);

        for (const [k, v] of this.collection[hashed]) if (k == key) return v;
        return null;
    };
    this.remove = function (key) {
        const hashed = hash(key);

        if (hashed in this.collection) {
            this.collection[hashed] = this.collection[hashed].filter((item) => item[0] !== key);
            if (this.collection[hashed].length === 0) delete this.collection[hashed];
        }
    };
};

const h = new HashTable();

h.add("09", 1);
h.add("18", 2);
h.add("27", 3);
h.add("36", 4);
console.log(h.collection[105]);
h.remove("18");
h.remove("27");
h.remove("09");
h.remove("36");
console.log(h.collection[105]);
