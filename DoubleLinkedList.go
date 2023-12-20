package list

import (
	"errors"
	"fmt"
)

// Node representa um nó na lista duplamente encadeada
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

// DoubleLinkedList é uma lista duplamente encadeada
type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// NewDoubleLinkedList cria uma nova lista duplamente encadeada
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// Add adiciona um valor ao final da lista
func (list *DoubleLinkedList) Add(value int) {
	newNode := &Node{Value: value, Prev: list.Tail, Next: nil}
	if list.Tail == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}
	list.Size++
}

// AddAtIndex adiciona um valor em uma posição específica na lista
func (list *DoubleLinkedList) AddAtIndex(index int, value int) {
	if index < 0 || index > list.Size {
		fmt.Println("Index out of bounds")
		return
	}

	newNode := &Node{Value: value, Prev: nil, Next: nil}
	if index == 0 {
		newNode.Next = list.Head
		if list.Head != nil {
			list.Head.Prev = newNode
		}
		list.Head = newNode
		if list.Tail == nil {
			list.Tail = newNode
		}
	} else if index == list.Size {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	} else {
		current := list.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		newNode.Prev = current
		newNode.Next = current.Next
		current.Next.Prev = newNode
		current.Next = newNode
	}

	list.Size++
}

// Remove remove o elemento na posição específica da lista
func (list *DoubleLinkedList) Remove(index int) {
	if index < 0 || index >= list.Size {
		fmt.Println("Index out of bounds")
		return
	}

	if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
	} else if index == 0 {
		list.Head = list.Head.Next
		list.Head.Prev = nil
	} else if index == list.Size-1 {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	} else {
		current := list.Head
		for i := 0; i < index; i++ {
			current = current.Next
		}
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}

	list.Size--
}

// Get retorna o valor na posição específica da lista
func (list *DoubleLinkedList) Get(index int) int {
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

// Set altera o valor na posição específica da lista
func (list *DoubleLinkedList) Set(index int, value int) {
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

// Size retorna o tamanho da lista
func (list *DoubleLinkedList) Size() int {
	return list.Size
}

func main() {
	var list IList = NewDoubleLinkedList()
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