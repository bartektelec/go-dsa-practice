package ds

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

func GetNth[T any](list *Node[T], index int) (*Node[T], bool) {
	current := *list
	for i := 0; i < index; i++ {
		if n := current.next; n != nil {
			current = *n
		} else {
			return nil, false
		}
	}

	return &current, true
}

func InsertNth[T any](list *Node[T], index int, value T) {
	before, _ := GetNth(list, index-1)
	after, _ := GetNth(list, index)

	current := Node[T]{
		value: value,
		prev:  before,
		next:  after,
	}

	before.next = &current
	after.prev = &current
}
