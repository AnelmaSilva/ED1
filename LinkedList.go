package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

type IList interface {
	Add(value int)
	AddAtIndex(index int, value int)
	Remove(index int)
	Get(index int) int
	Set(index int, value int)
	Size() int
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{Value: value, Next: nil}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.Size++
}

func (list *LinkedList) AddAtIndex(index int, value int) {
	if index < 0 || index > list.Size {
		fmt.Println("Index out of bounds")
		return
	}

	newNode := &Node{Value: value, Next: nil}
	if index == 0 {
		newNode.Next = list.Head
		list.Head = newNode
	} else {
		current := list.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		newNode.Next = current.Next
		current.Next = newNode
	}

	list.Size++
}

func (list *LinkedList) Remove(index int) {
	if index < 0 || index >= list.Size {
		fmt.Println("Index out of bounds")
		return
	}

	if index == 0 {
		list.Head = list.Head.Next
	} else {
		current := list.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}

	list.Size--
}

func (list *LinkedList) Get(index int) int {
	if index < 0 || index >= list.Size {
		fmt.Println("Index out of bounds")
		return -1
	}

	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value
}

func (list *LinkedList) Set(index int, value int) {
	if index < 0 || index >= list.Size {
		fmt.Println("Index out of bounds")
		return
	}

	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Value = value
}

func (list *LinkedList) Size() int {
	return list.Size
}

func main() {
	var list IList = &LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	fmt.Println("List size:", list.Size())
	fmt.Println("List elements:", list.Get(0), list.Get(1), list.Get(2))

	list.AddAtIndex(1, 4)
	fmt.Println("List elements after adding at index 1:", list.Get(0), list.Get(1), list.Get(2), list.Get(3))

	list.Remove(2)
	fmt.Println("List elements after removing at index 2:", list.Get(0), list.Get(1), list.Get(2))

	list.Set(1, 5)
	fmt.Println("List elements after setting value at index 1:", list.Get(0), list.Get(1), list.Get(2))
}