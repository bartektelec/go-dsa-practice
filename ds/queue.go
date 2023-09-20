package ds

type QueueNode[T any] struct {
	next  *QueueNode[T]
	value T
}

type Queue[T any] struct {
	head   *QueueNode[T]
	tail   *QueueNode[T]
	length int
}

func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Enqueue(value T) {
	node := &QueueNode[T]{
		value: value,
	}
	q.length++

	if q.Length() == 1 {
		q.head = node
		q.tail = q.head

		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() *QueueNode[T] {
	if q.head == nil {
		return nil
	}

	q.length--

	popped := q.head

	q.head = q.head.next
	popped.next = nil

	return popped
}

func (q *Queue[T]) Peek() T {
	return q.head.value
}
