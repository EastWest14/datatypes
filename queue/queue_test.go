package queue_test

import (
	. "datatypes/queue"
	"testing"
)

var (
	nilQueue        *Queue
	emptyQueue      *Queue
	oneElementQueue *Queue
	twoElementQueue *Queue
	veryLongQueue   *Queue
)

const LENGTH_OF_LONG_QUEUE = 100000

//Sets test variables to default values.
func setVariablesToDefaults() {
	nilQueue = nilQueue

	emptyQueue = NewQueue()

	oneElementQueue = NewQueue()
	oneElementQueue.Enqueue(0)

	twoElementQueue = NewQueue()
	twoElementQueue.Enqueue(0)
	twoElementQueue.Enqueue(1)

	veryLongQueue = NewQueue()
	for i := 0; i < LENGTH_OF_LONG_QUEUE; i++ {
		veryLongQueue.Enqueue(i)
	}
}

func TestNewQueue(t *testing.T) {
	newQueue := NewQueue()
	if newQueue == nil {
		t.Fatalf("Initialization of new queue fails")
	}
}
