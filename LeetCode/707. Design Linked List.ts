class LinkedListNode<T> {
    public next: LinkedListNode<T> | null = null;
    constructor(public data: T) {}
}

class MyLinkedList {
    private linkedList: LinkedListNode<number> | null = null;
    private length: number = 0;

    constructor() {}

    getNodeAtIndex(index: number): LinkedListNode<number> {
        let p = 0;
        let node = this.linkedList;
        while (p < index && node !== null) {
            node = node.next;
            p++;
        }
        return node;
    }

    get(index: number): number {
        const node = this.getNodeAtIndex(index);
        if (node === null) return -1;
        return node.data;
    }

    addAtHead(val: number): void {
        let newNode = new LinkedListNode<number>(val);
        newNode.next = this.linkedList;
        this.linkedList = newNode;
        this.length++;
    }

    addAtTail(val: number): void {
        let newNode = new LinkedListNode<number>(val);
        if (this.length === 0) {
            this.linkedList = newNode;
        } else {
            let tail = this.getNodeAtIndex(this.length - 1);
            tail.next = newNode;
        }
        this.length++;
    }

    addAtIndex(index: number, val: number): void {
        if (index >= this.length) return;

        if (index == 0) {
            this.addAtHead(val);
        } else {
            const nodeBeforeIndex = this.getNodeAtIndex(index - 1);
            const nodeAfterIndex = nodeBeforeIndex.next;
            const newNode = new LinkedListNode<number>(val);
            nodeBeforeIndex.next = newNode;
            newNode.next = nodeAfterIndex;
            this.length++;
        }
    }

    deleteAtIndex(index: number): void {
        if (index > this.length) return;

        if (index == 0) {
            this.linkedList = this.linkedList.next;
        } else {
            const nodeBeforeIndex = this.getNodeAtIndex(index - 1);
            const nodeAfterIndex = nodeBeforeIndex.next.next;
            nodeBeforeIndex.next = nodeAfterIndex;
        }
        this.length--;
    }

    toString(): string {
        let output = [];
        let node = this.linkedList;
        while (node !== null) {
            output.push(node.data);
            node = node.next;
        }
        return output.toString();
    }
}

var obj = new MyLinkedList();

// const cmd = ["MyLinkedList", "addAtHead", "get", "addAtHead", "addAtHead", "deleteAtIndex", "addAtHead", "get", "get", "get", "addAtHead", "deleteAtIndex"];
// const args = [[], [4], [1], [1], [5], [3], [7], [3], [3], [3], [1], [4]];

// const cmd = ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"];
// const args = [[], [1], [3], [1, 2], [1], [1], [1]];

const cmd = ["MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"];
const args = [[], [2], [1], [2], [7], [3], [2], [5], [5], [5], [6], [4]];

for (let i = 1; i < cmd.length; i++) {
    console.log("------------------------------");
    console.log("Command: ", cmd[i], ...args[i]);
    console.log("Output: ", obj[cmd[i]](...args[i]));
    console.log("Linked list: ", obj.toString());
}
