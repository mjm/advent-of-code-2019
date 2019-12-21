package day20

// Queue implements a simple FIFO queue using a circular buffer.
type Queue struct {
	items  []interface{}
	start  int
	length int
}

// NewQueue creates a new queue with a given initial size. If more elements
// are enqueued than the size, the queue's buffer will double in size.
func NewQueue(size int) *Queue {
	return &Queue{
		items:  make([]interface{}, size),
		start:  0,
		length: 0,
	}
}

// Empty checks if there are any elements in the queue.
func (q *Queue) Empty() bool {
	return q.length == 0
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(p interface{}) {
	q.growIfNeeded()
	i := q.nextIndex()
	q.items[i] = p
	q.length++
}

// Dequeue removes the element at the beginning of the queue and returns it.
// Returns false for the bool return value if the queue is empty.
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.length == 0 {
		return nil, false
	}

	p := q.items[q.start]
	q.start = (q.start + 1) % len(q.items)
	q.length--
	return p, true
}

func (q *Queue) growIfNeeded() {
	if q.length < len(q.items) {
		return
	}

	newItems := make([]interface{}, len(q.items)<<1)
	n := copy(newItems, q.items[q.start:])
	copy(newItems[n:], q.items[:q.start])
	q.items = newItems
	q.start = 0
}

func (q *Queue) nextIndex() int {
	return (q.start + q.length) % len(q.items)
}
