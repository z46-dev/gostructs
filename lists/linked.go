package lists

import (
	"fmt"

	"github.com/z46-dev/go-logger"
)

// Implements a doubly linked list in Go

type LinkedListNode[T any] struct {
	Value      T
	Next, Prev *LinkedListNode[T]
}

type LinkedList[T any] struct {
	Head, Tail *LinkedListNode[T]
	size       int
	comparator func(a, b T) bool
}

func NewLinkedList[T any](comparator func(a, b T) bool) *LinkedList[T] {
	return &LinkedList[T]{
		Head:       nil,
		Tail:       nil,
		size:       0,
		comparator: comparator,
	}
}

func (l *LinkedList[T]) Add(value T) {
	l.size++

	var node *LinkedListNode[T] = &LinkedListNode[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	node.Prev = l.Tail
	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList[T]) Remove(value T) bool {
	if l.Head == nil {
		return false
	}

	var node *LinkedListNode[T] = l.Head
	for node != nil {
		if l.comparator(node.Value, value) {
			l.size--

			if node.Prev != nil {
				node.Prev.Next = node.Next
			} else {
				l.Head = node.Next
			}

			if node.Next != nil {
				node.Next.Prev = node.Prev
			} else {
				l.Tail = node.Prev
			}

			return true
		}

		node = node.Next
	}

	return false
}

func (l *LinkedList[T]) Contains(value T) bool {
	if l.Head == nil {
		return false
	}

	var node *LinkedListNode[T] = l.Head
	for node != nil {
		if l.comparator(node.Value, value) {
			return true
		}
		node = node.Next
	}

	return false
}

func (l *LinkedList[T]) Get(index int) (zero T, found bool) {
	if index < 0 || index >= l.size {
		return zero, found
	}

	var node *LinkedListNode[T] = l.Head
	for range index {
		node = node.Next
	}

	return node.Value, !found
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.size = 0
}

func (l *LinkedList[T]) String() string {
	if l.Head == nil {
		return "[]"
	}

	var str string = "["
	var node *LinkedListNode[T] = l.Head
	for node != nil {
		str += fmt.Sprintf("%v", node.Value)
		node = node.Next
		if node != nil {
			str += ", "
		}
	}
	str += "]"

	return str
}

func LinkedListTests() {
	var list *LinkedList[int] = NewLinkedList[int](func(a, b int) bool {
		return a == b
	})

	var log *logger.Logger = logger.NewLogger().SetPrefix("[LinkedList]", logger.BoldRed)
	log.Basicf("Created an integer linked list\n")
	log.Basicf("List is empty: %v\n", list.IsEmpty())

	for i := range 10 {
		list.Add(i)
	}
}
