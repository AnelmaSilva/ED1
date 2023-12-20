package main

import (
	"errors"
	"fmt"
)

// IStack é a interface para uma pilha
type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

// Node representa um nó na lista ligada
type Node struct {
	Value int
	Next  *Node
}

// LinkedListStack é uma implementação de pilha usando uma lista ligada
type LinkedListStack struct {
	Top  *Node
	Size int
}

// Push adiciona um elemento ao topo da pilha
func (s *LinkedListStack) Push(value int) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
	s.Size++
}

// Pop remove e retorna o elemento no topo da pilha
func (s *LinkedListStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}

	value := s.Top.Value
	s.Top = s.Top.Next
	s.Size--

	return value, nil
}

// Peek retorna o elemento no topo da pilha sem removê-lo
func (s *LinkedListStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("a pilha está vazia")
	}

	return s.Top.Value, nil
}

// IsEmpty verifica se a pilha está vazia
func (s *LinkedListStack) IsEmpty() bool {
	return s.Size == 0
}

// Size retorna o tamanho da pilha
func (s *LinkedListStack) Size() int {
	return s.Size
}

func main() {
	stack := &LinkedListStack{}

	// Testando as operações da pilha
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Tamanho da pilha:", stack.Size())

	peekValue, err := stack.Peek()
	if err == nil {
		fmt.Println("Elemento no topo da pilha:", peekValue)
	} else {
		fmt.Println(err)
	}

	popValue, err := stack.Pop()
	if err == nil {
		fmt.Println("Elemento removido da pilha:", popValue)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Tamanho da pilha após a remoção:", stack.Size())
}
