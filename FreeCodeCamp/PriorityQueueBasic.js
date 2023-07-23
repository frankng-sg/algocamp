function PriorityQueue() {
    this.collection = [];
    this.printCollection = function () {
        console.log(this.collection);
    };
    // Only change code below this line
    this.dequeue = function () {
        return this.collection.shift()[0];
    };
    this.enqueue = function (val) {
        this.collection.push(val);
        let i = this.collection.length - 2;
        for (; i >= 0 && this.collection[i][1] > val[1]; i--) {
            this.collection[i + 1] = this.collection[i];
        }
        this.collection[i + 1] = val;
    };
    this.size = function () {
        return this.collection.length;
    };
    this.front = function () {
        return this.collection[0][0];
    };
    this.isEmpty = function () {
        return this.collection.length <= 0;
    };
    // Only change code above this line
}
const q = new PriorityQueue();
q.enqueue([1, 3]);
q.enqueue([1, 1]);
q.enqueue([2, 2]);
q.enqueue([1, 9]);
q.enqueue([2, 4]);
q.enqueue([2, 5]);
q.enqueue([3, 8]);
q.enqueue([3, 7]);
q.enqueue([3, 10]);

q.printCollection();
