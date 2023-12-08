package main

import (
	"fmt"
)

func binarySearch(list []string, search string) {
	var left, right, operations int

	left = 0
	right = len(list) - 1

	for left <= right {
		mid := (left + right) / 2
		current := list[mid]

		if search == current {
			fmt.Println("found:", current)
			break
		}

		if current > search {
			fmt.Println("smaller")
			right = mid - 1
		}

		if current < search {
			fmt.Println("bigger")
			left = mid + 1
		}

		operations++
	}

	fmt.Println("OPS n°: ", operations)
}

func main() {
	names := []string{
		"Alice", "Bob", "Charlie", "David", "Eva",
		"Frank", "Grace", "Hank", "Ivy", "Jack",
		"Karen", "Leo", "Mia", "Nathan", "Olivia",
		"Paul", "Quinn", "Rachel", "Sam", "Tracy",
		"Uma", "Victor", "Wendy", "Xander", "Yara",
		"Zane", "Bella", "Chris", "Diana", "Ethan",
		"Fiona", "George", "Holly", "Isaac", "Jade",
		"Kai", "Luna", "Mason", "Nora", "Oscar",
		"Penny", "Quincy", "Riley", "Sophia", "Tyler",
		"Ursula", "Vincent", "Willow", "Xavier", "Yasmine",
	}

	binarySearch(names, "Nathan")
}
