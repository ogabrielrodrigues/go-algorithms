package main

import (
	"fmt"

	"github.com/ogabrielrodrigues/go-algorithms/singly-linked-list/list"
)

type Vehicle struct {
	Brand string
	Name  string
	Year  string
	Price float64
}

func main() {
	list := list.List()
	list.Append("Value 1")
	list.Append("Value 2")
	list.Append("Value 3")
	list.Append("Value 4")

	list.Show()

	fmt.Printf("Value found on linked list: %s\n", list.Find("Value 3"))
}
