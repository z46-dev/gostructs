package lists

import "fmt"

type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

type Queue[T any] struct {
	front *QueueNode[T]
	rear  *QueueNode[T]
	size  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	node := &QueueNode[T]{Value: value}

	if q.size == 0 {
		q.front = node
		q.rear = node
	} else {
		q.rear.Next = node
		q.rear = node
	}

	q.size++
}

func (q *Queue[T]) Dequeue() *T {
	if q.size == 0 {
		return nil
	}

	value := q.front.Value
	q.front = q.front.Next
	q.size--

	if q.size == 0 {
		q.rear = nil
	}

	return &value
}

func (q *Queue[T]) Peek() *T {
	if q.size == 0 {
		return nil
	}

	return &q.front.Value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

func (q *Queue[T]) ToSlice() []T {
	slice := make([]T, 0, q.size)
	current := q.front

	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}

	return slice
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue: %v", q.ToSlice())
}

func QueueTest() {
	var q *Queue[int] = NewQueue[int]()

	for i := range 10 {
		q.Enqueue(i)
	}
}
