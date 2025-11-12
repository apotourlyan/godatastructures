package structures

import "errors"

// Compile-time interface verifications
var _ List[int] = &LinkedList[int]{}
var _ BasicList[int] = &BasicLinkedList[int]{}

// Represents a single node in a singly-linked list.
// Each node contains a value and a pointer to the next node.
type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
}

// Represents a singly-linked list for basic operations without comparison.
//
// This implementation provides fundamental list operations (add, remove, access)
// without requiring elements to be comparable.
//
// Design decisions:
//   - Head pointer: Enables O(1) access to first element
//   - Tail pointer: Enables O(1) AddLast and Last operations
//   - Size counter: Enables O(1) Size and IsEmpty operations
//   - No prev pointers: Keeps memory overhead low (not doubly-linked)
//   - No comparable constraint: Works with any type
//
// Space complexity: O(n) where n is the number of elements.
type BasicLinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

// Represents a singly-linked list implementation with head and tail pointers.
//
// Design decisions:
//   - Head pointer: Enables O(1) access to first element
//   - Tail pointer: Enables O(1) Add and Last operations
//   - Size counter: Enables O(1) Size and IsEmpty operations
//   - No prev pointers: Keeps memory overhead low (not a doubly-linked list)
//
// Space complexity: O(n) where n is the number of elements.
// Each node requires space for the value and one pointer.
type LinkedList[T comparable] struct {
	BasicLinkedList[T]
}

// Creates a new BasicLinkedList with optional initial values.
//
// Values are inserted in the order provided. If no values are given,
// an empty list is created.
//
// Time complexity: O(n) where n is the number of initial values.
//
// Example:
//
//	empty := NewBasicLinkedList[int]()
//	withValues := NewBasicLinkedList(1, 2, 3)
func NewBasicLinkedList[T any](values ...T) *BasicLinkedList[T] {
	l := &BasicLinkedList[T]{}
	size := len(values)
	if size == 0 {
		return l
	}

	// Use dummy node pattern to simplify construction
	dummy := &LinkedListNode[T]{}
	tail := dummy
	for _, v := range values {
		tail.Next = &LinkedListNode[T]{Value: v}
		tail = tail.Next
	}

	l.head = dummy.Next
	l.tail = tail
	l.size = size
	return l
}

// Creates a new LinkedList with optional initial values.
//
// Values are inserted in the order provided. If no values are given,
// an empty list is created.
//
// Time complexity: O(n) where n is the number of initial values.
//
// Example:
//
//	empty := NewLinkedList[int]()
//	withValues := NewLinkedList(1, 2, 3)
func NewLinkedList[T comparable](values ...T) *LinkedList[T] {
	basic := NewBasicLinkedList(values...)
	l := &LinkedList[T]{
		BasicLinkedList: *basic,
	}

	return l
}

// Prepends a value to the start of the list.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2)
//	l.AddFirst(0)  // List is now [0, 1, 2]
func (l *BasicLinkedList[T]) AddFirst(value T) {
	head := &LinkedListNode[T]{Value: value, Next: l.head}

	l.head = head
	if l.tail == nil {
		// Empty list: new node becomes both head and tail
		l.tail = head
	}

	l.size++
}

// Appends a value to the end of the list.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2)
//	l.AddLast(3)  // List is now [1, 2, 3]
func (l *BasicLinkedList[T]) AddLast(value T) {
	tail := &LinkedListNode[T]{Value: value}

	if l.head == nil {
		// Empty list: new node becomes both head and tail
		l.head = tail
		l.tail = tail
	} else {
		// Non-empty list: append to tail
		l.tail.Next = tail
		l.tail = tail
	}

	l.size++
}

// Removes a value from the start of the list.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.RemoveFirst()  // List is now [2, 3]
func (l *BasicLinkedList[T]) RemoveFirst() bool {
	if l.head == nil {
		return false
	}

	// Special case: one element in the list
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size--
		return true
	}

	head := l.head.Next
	l.head.Next = nil // Help GC
	l.head = head
	l.size--
	return true
}

// Removes a value from the end of the list.
//
// Time complexity: O(n)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.RemoveLast()  // List is now [1, 2]
func (l *BasicLinkedList[T]) RemoveLast() bool {
	if l.head == nil {
		return false
	}

	// Special case: one element in the list
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size--
		return true
	}

	node := l.head
	for node.Next != l.tail {
		node = node.Next
	}

	l.tail = node
	l.tail.Next = nil
	l.size--
	return true
}

// Returns the first element in the list.
//
// Returns ErrorEmptyList if the list is empty.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	first, _ := l.First()  // Returns 1
func (l *BasicLinkedList[T]) First() (T, error) {
	if l.head == nil {
		var zero T
		return zero, errors.New(ErrorEmptyList)
	}

	return l.head.Value, nil
}

// Returns the last element in the list.
//
// Returns ErrorEmptyList if the list is empty.
//
// Time complexity: O(1) - uses tail pointer
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	last, _ := l.Last()  // Returns 3
func (l *BasicLinkedList[T]) Last() (T, error) {
	if l.tail == nil {
		var zero T
		return zero, errors.New(ErrorEmptyList)
	}

	return l.tail.Value, nil
}

// Returns true if the list contains no elements.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList[int]()
//	l.IsEmpty()  // Returns true
//	l.Add(1)
//	l.IsEmpty()  // Returns false
func (l *BasicLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Size returns the number of elements in the list.
//
// Time complexity: O(1)
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.Size()  // Returns 3
func (l *BasicLinkedList[T]) Size() int {
	return l.size
}

// Inserts a value at the specified index.
//
// Valid indices are 0 to Size() inclusive. Index 0 inserts at the head,
// index Size() appends to the end (equivalent to Add).
//
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(n) where n is the index
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 3, 4)
//	l.InsertAt(1, 2)  // List is now [1, 2, 3, 4]
//	l.InsertAt(0, 0)  // List is now [0, 1, 2, 3, 4]
func (l *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > l.size {
		return errors.New(ErrorIndexOutOfRange)
	}

	// Special case: insert at head
	if index == 0 {
		l.head = &LinkedListNode[T]{Value: value, Next: l.head}
		if l.size == 0 {
			l.tail = l.head // Was empty, update tail
		}
		l.size++
		return nil
	}

	// Special case: insert at tail
	if index == l.size {
		l.tail.Next = &LinkedListNode[T]{Value: value}
		l.tail = l.tail.Next
		l.size++
		return nil
	}

	// Insert in middle: traverse to position
	prev := l.head
	for range index - 1 {
		prev = prev.Next
	}

	prev.Next = &LinkedListNode[T]{Value: value, Next: prev.Next}
	l.size++
	return nil
}

// Updates the element at the specified index.
//
// Valid indices are 0 to Size()-1.
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(n) where n is the index
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.UpdateAt(1, 4)  // Replaces 2 with 4, list is now [1, 4, 3]
func (l *LinkedList[T]) UpdateAt(index int, value T) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, errors.New(ErrorIndexOutOfRange)
	}

	node := l.head
	for range index {
		node = node.Next
	}

	old := node.Value
	node.Value = value
	return old, nil
}

// Removes the element at the specified index.
//
// Valid indices are 0 to Size()-1.
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(n) where n is the index
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.RemoveAt(1)  // Removes 2, list is now [1, 3]
func (l *LinkedList[T]) RemoveAt(index int) error {
	if index < 0 || index >= l.size {
		return errors.New(ErrorIndexOutOfRange)
	}

	// Special case: remove head
	if index == 0 {
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil // List becomes empty
		}
		l.size--
		return nil
	}

	// Remove from middle or end: traverse to position
	prev := l.head
	for range index - 1 {
		prev = prev.Next
	}

	target := prev.Next
	prev.Next = target.Next
	target.Next = nil // Help GC
	// Update tail if we removed the last element
	if target == l.tail {
		l.tail = prev
	}
	l.size--
	return nil
}

// Returns the element at the specified index.
//
// Valid indices are 0 to Size()-1.
// Returns ErrorIndexOutOfRange if index is invalid.
//
// Time complexity: O(n) where n is the index
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(10, 20, 30)
//	value, _ := l.GetAt(1)  // Returns 20
func (l *LinkedList[T]) GetAt(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, errors.New(ErrorIndexOutOfRange)
	}

	// Traverse to index
	node := l.head
	for range index {
		node = node.Next
	}

	return node.Value, nil
}

// Returns the index of the first occurrence of the specified value.
//
// Returns -1 if the value is not found.
//
// Time complexity: O(n) where n is the number of elements
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(10, 20, 30, 20)
//	index := l.IndexOf(20)  // Returns 1 (first occurrence)
//	index = l.IndexOf(99)   // Returns -1 (not found)
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

// Returns true if the list contains the specified value.
//
// Time complexity: O(n) where n is the number of elements
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3)
//	l.Contains(2)  // Returns true
//	l.Contains(9)  // Returns false
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

// Removes the first occurrence of the specified value.
//
// Returns true if the value was found and removed, false otherwise.
// The tail pointer is updated if the removed element was the last element.
//
// Time complexity: O(n) where n is the number of elements
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3, 2)
//	l.Remove(2)  // Removes first 2, list is now [1, 3, 2]
//	l.Remove(9)  // Returns false, list unchanged
func (l *LinkedList[T]) Remove(value T) bool {
	if l.head == nil {
		return false
	}

	// Special case: removing head
	if l.head.Value == value {
		if l.head == l.tail {
			l.tail = nil // List becomes empty
		}

		l.head = l.head.Next
		l.size--
		return true
	}

	// Search for value in rest of list
	prev := l.head
	for prev.Next != nil {
		if prev.Next.Value == value {
			target := prev.Next
			prev.Next = target.Next
			target.Next = nil // Help GC
			// Update tail if we removed the last element
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

// Replaces the first occurrence of the old value with the new value.
//
// Returns true if the value was found and updated, false otherwise.
//
// Time complexity: O(n) where n is the number of elements
//
// Space complexity: O(1)
//
// Example:
//
//	l := NewLinkedList(1, 2, 3, 2)
//	l.Update(2, 4)  // Updates first 2, list is now [1, 4, 3, 2]
//	l.Update(9, 3)  // Returns false, list unchanged
func (l *LinkedList[T]) Update(oldValue T, newValue T) bool {
	if l.head == nil {
		return false
	}

	node := l.head
	for node != nil {
		if node.Value == oldValue {
			node.Value = newValue
			return true
		}

		node = node.Next
	}

	return false
}
