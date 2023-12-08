package list

import "fmt"

// In this example 'value' attribute of node is string
type node struct {
	value string
	next  *node
}

type list struct {
	head *node
	tail *node
}

func List() *list {
	return &list{}
}

func (l *list) Append(value string) {
	node := &node{value: value}

	if l.head == nil {
		l.head = node
	}

	if l.tail != nil {
		l.tail.next = node
	}

	l.tail = node
}

func (l *list) Show() {
	node := l.head

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (l *list) Find(value string) string {
	node := l.head

	for node != nil {
		if node.value == value {
			break
		}

		node = node.next
	}

	if node != nil {
		return node.value
	}

	return ""
}
