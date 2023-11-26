package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{data: d, next: nil}

	if l.head == nil {
		l.head = list
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}

		ptr.next = list
	}
}

func Show(l *List) {
	ptr := l.head
	for ptr != nil {
		fmt.Printf("-> %v", ptr.data)
		ptr = ptr.next
	}
}

func (l *List) Popleft() (interface{}, bool) {
	if l.head == nil {
		return nil, false
	}

	poppedData := l.head.data
	l.head = l.head.next
	return poppedData, true
}
