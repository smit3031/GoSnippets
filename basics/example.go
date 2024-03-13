package basics

import "fmt"

func RunEx() {
	fmt.Println("Ok")

	myList := linkedList{}

	// Add nodes to the linked list
	myList.addNode(3)
	myList.addNode(7)
	myList.addNode(12)

	// Display the linked list
	myList.displayList()

}

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
}

func (ll *linkedList) addNode(val int) {
	newNode := &node{value: val, next: ll.head}
	ll.head = newNode
}

func (ll *linkedList) displayList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
