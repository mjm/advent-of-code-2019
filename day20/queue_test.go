package day20

import (
	"testing"

	"github.com/mjm/advent-of-code-2019/pkg/point"
	"github.com/stretchr/testify/assert"
)

func TestEmptyQueue(t *testing.T) {
	q := NewQueue(5)
	_, ok := q.Dequeue()
	assert.False(t, ok)
}

func TestQueueNoResize(t *testing.T) {
	q := NewQueue(10)
	q.Enqueue(newPoint(0, 0))
	q.Enqueue(newPoint(1, 0))

	p, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(0, 0), p)

	q.Enqueue(newPoint(0, 1))
	q.Enqueue(newPoint(1, 1))

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 0), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(0, 1), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 1), p)

	_, ok = q.Dequeue()
	assert.False(t, ok)
}

func TestQueueResizeAligned(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(newPoint(0, 0))
	q.Enqueue(newPoint(1, 0))
	q.Enqueue(newPoint(0, 1))

	// trigger a resize
	q.Enqueue(newPoint(1, 1))

	p, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(0, 0), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 0), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(0, 1), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 1), p)

	_, ok = q.Dequeue()
	assert.False(t, ok)
}

func TestQueueResizeNotAligned(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(newPoint(0, 0))
	q.Enqueue(newPoint(1, 0))
	q.Enqueue(newPoint(0, 1))
	q.Dequeue()
	q.Enqueue(newPoint(1, 1))

	// trigger a resize
	q.Enqueue(newPoint(2, 1))

	p, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 0), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(0, 1), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(1, 1), p)

	p, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, newPoint(2, 1), p)

	_, ok = q.Dequeue()
	assert.False(t, ok)
}

func newPoint(x, y int) point.Point2D {
	return point.Point2D{X: x, Y: y}
}
