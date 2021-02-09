package queue

const defaultInitialCapacity = 50

// Queue is backed by a slice used as a circular buffer.
type Queue struct {
	buf []interface{}
	// index of the first item
	start int
	// number of stored items
	len int
}

// New creates a new Queue. Optionally, custom initial capacity can be set.
func New(cap ...int) Queue {
	var _cap int
	if len(cap) > 0 {
		_cap = cap[0]
	} else {
		_cap = defaultInitialCapacity
	}
	return Queue{
		buf:   make([]interface{}, _cap),
		start: 0,
		len:   0,
	}
}

// extend the underlying slice and reset item order
func (q *Queue) extend() {
	// expects q.len == len(q.buf)
	newBuf := make([]interface{}, 2*len(q.buf))

	copy(newBuf, q.buf[q.start:])
	copy(newBuf[len(q.buf)-q.start:], q.buf[:q.start])

	q.start = 0
	q.buf = newBuf
}

// Push an item to the queue.
func (q *Queue) Push(item interface{}) {
	if q.len == len(q.buf) {
		q.extend()
	}
	q.buf[(q.start+q.len)%len(q.buf)] = item
	q.len++
}

// Pop removes and returns the first (oldest) item in the queue. If the queue
// is empty, ok is set to false, true otherwise.
func (q *Queue) Pop() (item interface{}, ok bool) {
	if q.len == 0 {
		return nil, false
	}
	item = q.buf[q.start]
	q.start = (q.start + 1) % len(q.buf)
	q.len--
	return item, true
}

// Front returns the first (about to be popped) item in the queue. If the queue
// is empty, ok is set to false, true otherwise.
func (q Queue) Front() (item interface{}, ok bool) {
	if q.len == 0 {
		return nil, false
	}
	return q.buf[q.start], true
}

// Len returns number of elements stored in the queue.
func (q Queue) Len() int {
	return q.len
}
