package deque

import "errors"

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type Deque[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// NewNode returns a new node.
func NewNode[T any](value T, next *Node[T], prev *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		next:  next,
		prev:  prev,
	}
}

// New returns a new deque.
func New[T any]() *Deque[T] {
	head, tail := NewNode[T](*new(T), nil, nil), NewNode[T](*new(T), nil, nil)
	head.next, tail.prev = tail, head
	return &Deque[T]{
		head: head,
		tail: tail,
	}
}

// IsEmpty returns true if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.head.next == d.tail && d.tail.prev == d.head
}

// PushFront adds a value to the front of the deque.
func (d *Deque[T]) PushFront(value T) {
	node := NewNode(value, d.head.next, d.head)
	node.next.prev, node.prev.next = node, node
}

// PushBack adds a value to the back of the deque.
func (d *Deque[T]) PushBack(value T) {
	node := NewNode(value, d.tail, d.tail.prev)
	node.next.prev, node.prev.next = node, node
}

// PopFront removes a value at the front of the deque.
func (d *Deque[T]) PopFront() (T, error) {
	if d.IsEmpty() {
		return *new(T), errors.New("deque is empty")
	}
	res := d.head.next
	res.prev.next, res.next.prev = res.next, res.prev
	res.next, res.prev = nil, nil
	return res.value, nil
}

// PopBack removes a value at the back of the deque.
func (d *Deque[T]) PopBack() (T, error) {
	if d.IsEmpty() {
		return *new(T), errors.New("deque is empty")
	}
	res := d.tail.prev
	res.prev.next, res.next.prev = res.next, res.prev
	res.next, res.prev = nil, nil
	return res.value, nil
}

// Front returns the value at the front of the deque.
func (d *Deque[T]) Front() (T, error) {
	if d.IsEmpty() {
		return *new(T), errors.New("deque id empty")
	}
	return d.head.next.value, nil
}

// Back returns the value at the back of the deque.
func (d *Deque[T]) Back() (T, error) {
	if d.IsEmpty() {
		return *new(T), errors.New("deque id empty")
	}
	return d.tail.prev.value, nil
}

// Len returns the length of the deque.
// O(n) time complexity so should be used sparingly.
func (d *Deque[T]) Len() int {
	length := 0
	for cur := d.head.next; cur != d.tail; cur = cur.next {
		length++
	}
	return length
}

// Get returns the value at the given index.
// O(n) time complexity so should be used sparingly.
func (d *Deque[T]) Get(index int) (T, error) {
	if index < 0 || index >= d.Len() {
		return *new(T), errors.New("index out of bounds")
	}
	cur := d.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.value, nil
}
