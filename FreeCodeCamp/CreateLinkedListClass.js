function LinkedList() {
    var length = 0;
    var head = null;

    var Node = function (element) {
        this.element = element;
        this.next = null;
    };

    this.head = function () {
        return head;
    };

    this.size = function () {
        return length;
    };

    this.add = function (element) {
        var newNode = new Node(element);
        if (!head) head = newNode;
        else {
            let node = head;
            while (node.next !== null) node = node.next;
            node.next = newNode;
        }
        length++;
    };
}

let ll = new LinkedList();
ll.add(1);
ll.add(2);
ll.add(3);
console.log(ll.head().element);
console.log(ll.head().next);
