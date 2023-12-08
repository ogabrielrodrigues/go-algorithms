package main

import (
	"fmt"
)

type vehicle struct {
	name  string
	brand string
	year  string
	price uint64
}

type node struct {
	value vehicle
	next  *node
}

type list struct {
	head *node
	tail *node
}

func List() *list {
	return &list{}
}

func (l *list) append(value vehicle) {
	node := &node{value: value}

	if l.head == nil {
		l.head = node
	}

	if l.tail != nil {
		l.tail.next = node
	}

	l.tail = node
}

func (l *list) show() {
	node := l.head

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (l *list) find(name string) *vehicle {
	node := l.head

	for node != nil {
		if node.value.name == name {
			break
		}

		node = node.next
	}

	if node != nil {
		return &node.value
	}

	return nil
}

func main() {
	list := list{}
	list.append(vehicle{name: "Jetta", brand: "Volkswagen", year: "2012", price: 16000})
	list.append(vehicle{name: "RS3", brand: "Audi", year: "2018", price: 54000})
	list.append(vehicle{name: "Civic", brand: "Honda", year: "2015", price: 23000})
	list.append(vehicle{name: "SL65 AMG", brand: "Mercedes Benz", year: "2009", price: 315000})

	fmt.Printf("Value found on linked list: %v\n", list.find("Civic"))
}
