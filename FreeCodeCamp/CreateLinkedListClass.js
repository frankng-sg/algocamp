function LinkedList() {
    var length = 0;
    var head = null;

    var Node = function (element) {
        this.element = element;
        this.next = null;
    };

    this.size = function () {
        return length;
    };

    this.head = function () {
        return head;
    };

    this.add = function (element) {
        var node = new Node(element);
        if (head === null) {
            head = node;
        } else {
            var currentNode = head;

            while (currentNode.next) {
                currentNode = currentNode.next;
            }

            currentNode.next = node;
        }

        length++;
    };

    this.remove = function (element) {
        if (!head) return;
        if (head.element === element) {
            head = head.next;
            length--;
            return;
        }
        let prev = head;
        let curr = head.next;
        while (curr) {
            if (curr.element === element) {
                prev.next = curr.next;
                length--;
                break;
            }
            prev = curr;
            curr = curr.next;
        }
    };
}

let ll = new LinkedList();
ll.add(1);
console.log(ll.size());
ll.remove(1);
ll.remove(2);
console.log(ll.size());
