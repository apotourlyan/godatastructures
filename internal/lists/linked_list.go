package linkedlists

import "errors"

type LinkedListNode[T comparable] struct {
	Value T
	Next  *LinkedListNode[T]
}

type LinkedList[T comparable] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

func NewLinkedList[T comparable](values ...T) *LinkedList[T] {
	size := len(values)
	if size == 0 {
		return &LinkedList[T]{size: size}
	}

	dummy := &LinkedListNode[T]{}
	node := dummy
	for _, v := range values {
		node.Next = &LinkedListNode[T]{Value: v}
		node = node.Next
	}

	return &LinkedList[T]{head: dummy.Next, tail: node, size: len(values)}
}

func (l *LinkedList[T]) Add(value T) {
	tail := &LinkedListNode[T]{Value: value}

	if l.head == nil {
		l.head = tail
		l.tail = tail
	} else {
		l.tail.Next = tail
		l.tail = tail
	}

	l.size++
}

func (l *LinkedList[T]) Remove(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		if l.head == l.tail {
			l.tail = nil
		}

		l.head = l.head.Next
		l.size--
		return true
	}

	prev := l.head
	for prev.Next != nil {
		if prev.Next.Value == value {
			target := prev.Next
			prev.Next = target.Next
			target.Next = nil
			if target == l.tail {
				l.tail = prev
			}
			l.size--
			return true
		}

		prev = prev.Next
	}

	return false
}

func (l *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > l.size {
		return errors.New(ErrorIndexOutOfRange)
	}

	if index == 0 {
		l.head = &LinkedListNode[T]{Value: value, Next: l.head}
		if l.size == 0 {
			l.tail = l.head
		}
		l.size++
		return nil
	}

	if index == l.size {
		l.Add(value)
		return nil
	}

	prev := l.head
	for range index - 1 {
		prev = prev.Next
	}

	prev.Next = &LinkedListNode[T]{Value: value, Next: prev.Next}
	l.size++
	return nil
}

func (l *LinkedList[T]) RemoveAt(index int) error {
	if index < 0 || index >= l.size {
		return errors.New(ErrorIndexOutOfRange)
	}

	if index == 0 {
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return nil
	}

	prev := l.head
	for range index - 1 {
		prev = prev.Next
	}

	target := prev.Next
	prev.Next = target.Next
	target.Next = nil
	if target == l.tail {
		l.tail = prev
	}
	l.size--
	return nil
}

func (l *LinkedList[T]) GetAt(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, errors.New(ErrorIndexOutOfRange)
	}

	node := l.head
	for range index {
		node = node.Next
	}

	return node.Value, nil
}

func (l *LinkedList[T]) IndexOf(value T) int {
	node := l.head
	for i := 0; node != nil; i++ {
		if node.Value == value {
			return i
		}

		node = node.Next
	}

	return -1
}

func (l *LinkedList[T]) Contains(value T) bool {
	node := l.head

	for node != nil {
		if node.Value == value {
			return true
		}

		node = node.Next
	}

	return false
}

func (l *LinkedList[T]) First() (T, error) {
	if l.head == nil {
		var zero T
		return zero, errors.New(ErrorEmptyList)
	}

	return l.head.Value, nil
}

func (l *LinkedList[T]) Last() (T, error) {
	if l.tail == nil {
		var zero T
		return zero, errors.New(ErrorEmptyList)
	}

	return l.tail.Value, nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
