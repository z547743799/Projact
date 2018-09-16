package queue

type node struct {
	data interface{}
	next *node
}

type Queue struct {
	head  *node
	tail  *node
	count int
}

func New() *Queue {
	q := &Queue{}
	return q
}

func (q *Queue) Len() int {
	return q.count
}

func (q *Queue) Enqueue(item interface{}) {
	n := &node{data: item}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--

	return n.data
}

func (q *Queue) Peek() interface{} {
	n := q.head
	if n == nil || n.data == nil {
		return nil
	}
	return n.data
}
