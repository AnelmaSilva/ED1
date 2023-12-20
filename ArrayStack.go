package stack

import "errors"

type ArrayStack struct {
	values []int
	size   int
}

func (stack *ArrayStack) Init(size int) error {
	if size > 0 {
		stack.values = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayStack with size <= 0")
	}
}

/**
 * Duplica o vetor.
 */
func (stack *ArrayStack) doubleArray() {
	curSize := len(stack.values)
	doubledValues := make([]int, 2*curSize)

	for i := 0; i < curSize; i++ {
		doubledValues[i] = stack.values[i]
	}
	stack.values = doubledValues
}

func (stack *ArrayStack) Push(val int) {
	//verificar se array está lotado
	if stack.size >= len(stack.values) {
		stack.doubleArray()
	}
	stack.values[stack.size] = val
	stack.size++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.size > 0 {
		stack.size--
		return stack.values[stack.size], nil
	} else {
		return -1, errors.New("Can't pop from empty stack")
	}
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.size > 0 {
		return stack.values[stack.size-1], nil
	} else {
		return -1, errors.New("Can't get the peek of an empty stack")
	}
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *ArrayStack) Size() int {
	return stack.size
}
