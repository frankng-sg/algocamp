class CircularQueue {
    constructor(size) {
        this.queue = [];
        this.read = 0;
        this.write = 0;
        this.max = size - 1;

        while (size > 0) {
            this.queue.push(null);
            size--;
        }
    }

    print() {
        return this.queue;
    }

    enqueue(item) {
        // Only change code below this line
        let next = this.write;
        if (this.queue[next] !== null) return null;
        this.queue[next] = item;
        if (++next > this.max) next = next - this.max - 1;
        this.write = next;
        return item;
        // Only change code above this line
    }

    dequeue() {
        // Only change code below this line
        if (this.queue[this.read] == null) return null;
        const val = this.queue[this.read];
        this.queue[this.read] = null;
        if (++this.read > this.max) this.read = this.read - this.max - 1;
        return val;
        // Only change code above this line
    }
}

let q = new CircularQueue(5);

q.enqueue(1);
console.log(q.print());
q.enqueue(2);
console.log(q.print());
q.enqueue(3);
console.log(q.print());
q.enqueue(4);
console.log(q.print());
q.enqueue(5);
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.enqueue(6);
console.log(q.print());
q.enqueue(7);
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.enqueue(8);
console.log(q.print());
q.enqueue(9);
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.enqueue(10);
console.log(q.print());
q.enqueue(11);
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
q.dequeue();
console.log(q.print());
