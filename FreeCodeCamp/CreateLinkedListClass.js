function LinkedList() {
    let length = 0;
    let head = null;

    let Node = function (element) {
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
        let node = new Node(element);
        if (head === null) {
            head = node;
        } else {
            let currentNode = head;

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
    this.isEmpty = function () {
        return length <= 0;
    };
    this.indexOf = function (element) {
        let idx = 0;
        let curr = head;
        while (curr && curr.element != element) {
            idx++;
            curr = curr.next;
        }
        if (!curr) return -1;
        return idx;
    };
    this.elementAt = function (idx) {
        if (idx < 0 || idx >= length) return undefined;
        let curr = head;
        while (curr && idx-- > 0) curr = curr.next;
        return curr.element;
    };
    this.removeAt = function (idx) {
        if (idx < 0 || idx >= length) return null;
        let val;
        if (idx === 0) {
            val = head.element;
            head = head.next;
        } else {
            let prev = head;
            let curr = head;
            while (curr && idx-- > 0) {
                prev = curr;
                curr = curr.next;
            }
            prev.next = curr.next;
            val = curr.element;
        }
        length--;
        return val;
    };
    this.addAt = function (idx, element) {
        if (idx < 0 || idx >= length) return false;
        let newNode = new Node(element);
        if (idx == 0) {
            newNode.next = head;
            head = newNode;
        } else {
            let prev = head;
            let curr = head;
            while (idx-- > 0) {
                prev = curr;
                curr = curr.next;
            }
            prev.next = newNode;
            newNode.next = curr;
        }
        length++;
        return true;
    };
}
