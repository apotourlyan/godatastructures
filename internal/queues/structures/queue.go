package structures

const ErrorEmptyQueue = "queue is empty"

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}
