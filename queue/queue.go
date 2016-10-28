package queue

type Queue struct {
}

func NewQueue() *Queue {
	return nil
}

func (q *Queue) Length() int {
	return 0
}

func (q *Queue) Peek() (value interface{}) {
	return nil
}

func (q *Queue) Enqueue(value interface{}) {
}

func (q *Queue) Dequeue() (valueRemoved interface{}) {
	return nil
}
