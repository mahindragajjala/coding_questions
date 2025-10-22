package main

import "fmt"

// Node represents each element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the entire list
type LinkedList struct {
	head *Node
}

// InsertAtEnd adds a new node at the end
func (list *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{data: value}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// InsertAtBeginning adds a node at the start
func (list *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{data: value, next: list.head}
	list.head = newNode
}

// InsertAfter inserts a node after a given value
func (list *LinkedList) InsertAfter(afterValue int, newValue int) {
	current := list.head
	for current != nil && current.data != afterValue {
		current = current.next
	}

	if current == nil {
		fmt.Println("Value not found in list!")
		return
	}

	newNode := &Node{data: newValue, next: current.next}
	current.next = newNode
}

// DeleteByValue removes a node with a given value
func (list *LinkedList) DeleteByValue(value int) {
	if list.head == nil {
		fmt.Println("List is empty!")
		return
	}

	// If head is to be deleted
	if list.head.data == value {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		fmt.Println("Value not found!")
		return
	}

	current.next = current.next.next
}

// Search finds if a value exists in the list
func (list *LinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

// Reverse reverses the linked list
func (list *LinkedList) Reverse() {
	var prev *Node
	current := list.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

// Display prints the list
func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty!")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Length returns the total number of nodes
func (list *LinkedList) Length() int {
	count := 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func main() {
	list := LinkedList{}

	// Insert elements
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtBeginning(5)
	list.InsertAfter(20, 25)

	fmt.Print("Initial list: ")
	list.Display()

	// Delete element
	list.DeleteByValue(10)
	fmt.Print("After deleting 10: ")
	list.Display()

	// Search element
	fmt.Println("Is 25 present?", list.Search(25))
	fmt.Println("Is 100 present?", list.Search(100))

	// Reverse list
	list.Reverse()
	fmt.Print("Reversed list: ")
	list.Display()

	// Print length
	fmt.Println("Length of list:", list.Length())
}
