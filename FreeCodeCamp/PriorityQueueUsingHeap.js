function PriorityQueue() {
    this.collection = [];
    this.printCollection = function () {
        console.log(this.collection);
    };
    // Only change code below this line
    this.leftChild = function (parentIndex) {
        return parentIndex * 2 + 1;
    };
    this.rightChild = function (parentIndex) {
        return parentIndex * 2 + 2;
    };
    this.swap = function (i, j) {
        let tmp = this.collection[i];
        this.collection[i] = this.collection[j];
        this.collection[j] = tmp;
    };
    this.heapifyDown = function (parentIndex) {
        const l = this.leftChild(parentIndex);
        const r = this.rightChild(parentIndex);
        let smallest = parentIndex;
        if (l < this.collection.length && this.collection[l][1] < this.collection[smallest][1]) smallest = l;
        if (r < this.collection.length && this.collection[r][1] < this.collection[smallest][1]) smallest = r;
        if (smallest !== parentIndex) {
            this.swap(smallest, parentIndex);
            this.heapifyDown(smallest);
        }
    };
    this.heapifyUp = function (childIndex) {
        const parentIndex = (childIndex - 1) >> 1;
        if (parentIndex >= 0 && this.collection[parentIndex][1] > this.collection[childIndex][1]) {
            this.swap(parentIndex, childIndex);
            this.heapifyUp(parentIndex);
        }
    };
    this.dequeue = function () {
        const val = this.collection.shift();
        this.heapifyDown(0);
        return val;
    };
    this.enqueue = function (val) {
        this.collection.push(val);
        this.heapifyUp(this.collection.length - 1);
    };
    this.size = function () {
        return this.collection.length;
    };
    this.front = function () {
        return this.collection[0];
    };
    this.isEmpty = function () {
        return this.collection.length <= 0;
    };
    // Only change code above this line
}
