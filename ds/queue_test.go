package ds

import "testing"

func TestQueueShouldEnqueueNodeAndPeek(t *testing.T) {
	q := Queue[int]{}

	if q.Length() != 0 {
		t.Errorf("Init queue length was wrong")
	}

	q.Enqueue(10)

	if q.Length() != 1 {
		t.Errorf("Length was %v but expected 1", q.Length())
	}

	if q.Peek() != 10 {
		t.Errorf("Peek didnt resolve the enqueued item")
	}
}

func TestQueueShouldEnqueueMultipleNodesAndDeque(t *testing.T) {
	q := Queue[int]{}

	if q.Length() != 0 {
		t.Errorf("Init queue length was wrong")
	}

	enqueuedValues := [3]int{10, 5, 3}

	for _, value := range enqueuedValues {
		q.Enqueue(value)
	}

	for _, expected := range enqueuedValues {

		dequed := q.Deque()
		if dequed.value != expected {
			t.Errorf("Dequeued value was %v but expected %v", dequed.value, expected)
		}
	}

	if q.Deque() != nil {
		t.Errorf("Overflowed value expected to be nil")
	}
}
