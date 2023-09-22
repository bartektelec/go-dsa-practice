package ds

import "math"

type StackNode[T any] struct {
	value T
	prev  *StackNode[T]
}

type Stack[T any] struct {
	head   *StackNode[T]
	tail   *StackNode[T]
	length int
}

func (this *Stack[T]) Length(item T) {
}

func (this *Stack[T]) Push(item T) {
	node := StackNode[T]{
		value: item,
	}

	this.length++

	if this.head == nil {
		this.head = &node
		return
	}

	node.prev = this.head
	this.head = &node
}

func (this *Stack[T]) Pop() *T {
	this.length = int(math.Max(float64(0), float64(this.length-1)))

	h := this.head

	if h == nil {
		return nil
	}

	if this.length == 0 {
		this.head = nil
	} else {
		this.head = h.prev
	}

	return &h.value
}

func (this *Stack[T]) Peek() *T {
	if this.head == nil {
		return nil
	}
	return &this.head.value
}
