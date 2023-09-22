package ds

import "testing"

func TestStackShouldPeek(t *testing.T) {
	stack := Stack[int]{
		head: &StackNode[int]{
			value: 5,
			prev:  nil,
		},
	}

	if result := stack.Peek(); *result != 5 {
		t.Error("Peek didnt return 5")
	}
}

func TestStackShouldPeekIfEmpty(t *testing.T) {
	stack := Stack[int]{}

	if result := stack.Peek(); result != nil {
		t.Error("Peek didnt return nil")
	}
}

func TestStack(t *testing.T) {
	stack := Stack[int]{}

	expected := []int{30, 20, 10}

	for _, v := range expected {
		stack.Push(v)
	}

	for i := len(expected) - 1; i >= 0; i-- {
		exp_value := expected[i]

		if result := stack.Pop(); *result != exp_value {
			t.Errorf("Popped value was %v but expected %v", *result, exp_value)
		}

	}
}
