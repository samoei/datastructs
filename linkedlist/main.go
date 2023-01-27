package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(data *node) {
	second := l.head
	l.head = data
	l.head.next = second
	l.length++
}

func (l LinkedList) print() {
	current := l.head
	for l.length > 0 {
		fmt.Printf("%d->", current.data)
		current = current.next
		l.length--
	}
	fmt.Println('\n')

}

func (l *LinkedList) delete(data int) {
	prev := l.head
	current := l.head

	for current.next != nil {
		next := current.next
		if current.data == data {
			current.next = nil
			prev.next = next
			l.length--
		} else {
			prev = current
			current = next
			next = current.next
		}
	}

}

func main() {
	myList := LinkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 40}
	node5 := &node{data: 50}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	myList.print()

	myList.delete(50)
	myList.print()

}
