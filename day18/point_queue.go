package day18

import "github.com/mjm/advent-of-code-2019/pkg/point"

// PointQueue implements a simple FIFO queue for points using a circular buffer.
type PointQueue struct {
	items  []point.Point2D
	start  int
	length int
}

// NewPointQueue creates a new queue with a given initial size. If more points
// are enqueued than the size, the queue's buffer will double in size.
func NewPointQueue(size int) *PointQueue {
	return &PointQueue{
		items:  make([]point.Point2D, size),
		start:  0,
		length: 0,
	}
}

// Empty checks if there are any points in the queue.
func (q *PointQueue) Empty() bool {
	return q.length == 0
}

// Enqueue adds a point to the end of the queue.
func (q *PointQueue) Enqueue(p point.Point2D) {
	q.growIfNeeded()
	i := q.nextIndex()
	q.items[i] = p
	q.length++
}

// Dequeue removes the point at the beginning of the queue and returns it.
// Returns false for the bool return value if the queue is empty.
func (q *PointQueue) Dequeue() (point.Point2D, bool) {
	if q.length == 0 {
		return point.Point2D{}, false
	}

	p := q.items[q.start]
	q.start = (q.start + 1) % len(q.items)
	q.length--
	return p, true
}

func (q *PointQueue) growIfNeeded() {
	if q.length < len(q.items) {
		return
	}

	newItems := make([]point.Point2D, len(q.items)<<1)
	n := copy(newItems, q.items[q.start:])
	copy(newItems[n:], q.items[:q.start])
	q.items = newItems
	q.start = 0
}

func (q *PointQueue) nextIndex() int {
	return (q.start + q.length) % len(q.items)
}
