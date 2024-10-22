class Node {
    constructor(value = null, next = null, prev = null) {
        this.value = value;
        this.next = next;
        this.prev = prev;
    }
}

class circularDoublyLinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
    }

    // Utility function to add node(for testing purposes)
    append(value) {
        const newNode = new Node(value);
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            newNode.next = newNode;
            newNode.prev = newNode;
        } else {
        newNode.prev = this.tail;
        newNode.next = this.head;
        this.tail.next = newNode;
        this.tail = newNode;
    }
}

deleteAt(index) {
    if (!this.head || index < 0) return;

    let current = this.head;
    let count = 0;
    const size = this.getSize(); // Utlitly fuction to get the size of the list

    if (index >= size) return; // Prevent out-of-bounds access

    // Special case: deleting the head node
    if (index === 0) {
        if (this.head === this.tail) {
            // Only one node in the list
            this.head = null;
            this.tail = null;     
        } else {
            this.tail.next = this.head.next;
            this.head.next.prev = this.tail;
            this.head = this.head.next;
        }
        return;
    }

    // Traverse the list to find the node at the given index 
    while (count < index) {
        current = current.next;
        count++;
    }

    // Now 'current' is the node to delete
    current.prev.next = current.next;
    current.next.prev = current.prev;

    // Special case: delteting the tail node
    if (current === this.tial) {
        this.tail = current.prev
    }
}

// Utility function to get the size of the list 
getSize() {
    if (!this.head) return 0;
    let count = 1;
    let current = this.head;
    while (current.next !== this.head) {
        count++;
        current = current.next;
    }
    return count;
}

// Utility function to display the list (for testing purposes)
printList() {
    if (!this.head) {
        console.log("Empty List");
        return;
    }
    let result = [];
    let current = this.head;
    do {
        result.push(current.value);
        current = current.next;
    } while (current !== this.head);
    console.log(result.join(" <-> "));
    }
}

// Example Usage
const list = new circularDoublyLinkedList();
list.append(1);
list.append(2);
list.append(3);

console.log("Before deletion:");
list.printList(); // 1 <-> 2 <-> 3 <-> 1 (circular)

list.deleteAt(1);

console.log("After deletion:");
list.printList(); // 1 <-> 3 <-> 1 (circular)

module.exports = { Node, CircularDoublyLinkedList };