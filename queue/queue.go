package queue

type Queue struct {
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Length() int {
	return 0
}

func (q *Queue) Peek() (value interface{}, err error) {
	return nil, nil
}

func (q *Queue) Enqueue(value interface{}) {
}

func (q *Queue) Dequeue() (valueRemoved interface{}, err error) {
	return nil, nil
}
