package linkedlists

const ErrorEmptyList = "list is empty"
const ErrorIndexOutOfRange = "index is out of the range of possible values"

type List[T comparable] interface {
	Add(value T)
	Remove(value T) bool
	InsertAt(index int, value T) error
	RemoveAt(index int) error
	GetAt(index int) (T, error)
	IndexOf(value T) int
	Contains(value T) bool
	First() (T, error)
	Last() (T, error)
	IsEmpty() bool
	Size() int
}
