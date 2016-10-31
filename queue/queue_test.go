package queue_test

import (
	. "datatypes/queue"
	"sync"
	"testing"
	"time"
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
	twoElementQueue.Enqueue("0")
	twoElementQueue.Enqueue("1")

	veryLongQueue = NewQueue()
	for i := 0; i < LENGTH_OF_LONG_QUEUE; i++ {
		veryLongQueue.Enqueue(i)
	}
}

//*************** Public Interface Test ***************

func TestNewQueue(t *testing.T) {
	newQueue := NewQueue()
	if newQueue == nil {
		t.Fatalf("Initialization of new queue fails")
	}
	if newQueue.Length() != 0 {
		t.Error("Queue initialized to non-zero value")
	}
}

func TestLength(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		queueInstance  *Queue
		expectedLength int
	}{
		{queueInstance: nilQueue, expectedLength: 0},
		{queueInstance: emptyQueue, expectedLength: 0},
		{queueInstance: oneElementQueue, expectedLength: 1},
		{queueInstance: twoElementQueue, expectedLength: 2},
		{queueInstance: veryLongQueue, expectedLength: LENGTH_OF_LONG_QUEUE},
	}

	for i, aCase := range cases {
		length := aCase.queueInstance.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected queue length %d, got %d", i, aCase.expectedLength, length)
		}
	}
}

func TestPeek(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		queueInstance *Queue
		expectedValue interface{}
		expectError   bool
	}{
		{queueInstance: emptyQueue, expectedValue: nil, expectError: true},
		{queueInstance: oneElementQueue, expectedValue: 0, expectError: false},
		//Check it stays the same
		{queueInstance: oneElementQueue, expectedValue: 0, expectError: false},
		{queueInstance: twoElementQueue, expectedValue: "0", expectError: false},
		{queueInstance: veryLongQueue, expectedValue: 0, expectError: false},
	}

	for i, aCase := range cases {
		peekValue, err := aCase.queueInstance.Peek()

		if !aCase.expectError {
			if err != nil {
				t.Errorf("Error in case %d. Expected no error, got %s", i, err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("Error in case %d. Expected error, got no error", i)
			}
		}
		if peekValue != aCase.expectedValue {
			t.Errorf("Error in case %d. Expected value %v, got %v", i, aCase.expectedValue, peekValue)
		}
	}
}

type dequeueExpectation struct {
	value       interface{}
	expectError bool
}

func TestEnqueue(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		queueInstance              *Queue
		enqueueSequence            []interface{}
		expectedLengthAfterEnqueue int
		dequeueExpectations        []dequeueExpectation
	}{
		{emptyQueue, []interface{}{0}, 1, []dequeueExpectation{{value: 0, expectError: false}, {value: nil, expectError: true}}},
		{oneElementQueue, []interface{}{1, 2}, 3, []dequeueExpectation{{value: 0, expectError: false}, {value: 1, expectError: false}, {value: 2, expectError: false}, {value: nil, expectError: true}}},
		//Test enqueue after a dequeue
		{oneElementQueue, []interface{}{"New only element"}, 1, []dequeueExpectation{{value: "New only element", expectError: false}, {value: nil, expectError: true}}},
		{veryLongQueue, []interface{}{"last value"}, LENGTH_OF_LONG_QUEUE + 1, []dequeueExpectation{{value: 0, expectError: false}}},
	}

	for caseNumber, aCase := range cases {
		//Enqueue all the values
		for _, enqueuValue := range aCase.enqueueSequence {
			aCase.queueInstance.Enqueue(enqueuValue)
		}

		length := aCase.queueInstance.Length()
		if length != aCase.expectedLengthAfterEnqueue {
			t.Errorf("Error in case %d. Expected value %v, got %v", caseNumber, aCase.expectedLengthAfterEnqueue, length)
		}

		for i, dequeueExpect := range aCase.dequeueExpectations {
			dequeuValue, err := aCase.queueInstance.Dequeue()

			//Check error value is expected
			if !dequeueExpect.expectError {
				if err != nil {
					t.Errorf("Error in case %d, dequeu %d. Expected no error, got %s", caseNumber, i, err.Error)
				}
			} else {
				if err == nil {
					t.Errorf("Error in case %d, dequeu %d. Expected an error, got no error", caseNumber, i)
				}
			}

			//Check dequeued value is correct
			if dequeuValue != dequeueExpect.value {
				t.Errorf("Error in case %d, dequeu %d. Expected value %v, got %v", caseNumber, i, dequeueExpect.value, dequeuValue)
			}
		}
	}

	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("Enqueue on a nil queue should cause a panic, did not")
		}
	}()
	nilQueue.Enqueue(0)
}

func TestRemove(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		queueInstance       *Queue
		dequeueExpectations []dequeueExpectation
		expectedLength      int
	}{
		{emptyQueue, []dequeueExpectation{{value: nil, expectError: true}}, 0},
		//Try double remove on empty queue
		{oneElementQueue, []dequeueExpectation{{value: 0, expectError: false}, {value: nil, expectError: true}, {value: nil, expectError: true}}, 0},
		{twoElementQueue, []dequeueExpectation{{value: "0", expectError: false}}, 1},
	}

	for caseNumber, aCase := range cases {
		for i, dequeueExpect := range aCase.dequeueExpectations {
			dequeuValue, err := aCase.queueInstance.Dequeue()

			//Check error value is expected
			if !dequeueExpect.expectError {
				if err != nil {
					t.Errorf("Error in case %d, dequeu %d. Expected no error, got %s", caseNumber, i, err.Error)
				}
			} else {
				if err == nil {
					t.Errorf("Error in case %d, dequeu %d. Expected an error, got no error", caseNumber, i)
				}
			}

			//Check dequeued value is correct
			if dequeuValue != dequeueExpect.value {
				t.Errorf("Error in case %d, dequeu %d. Expected value %v, got %v", caseNumber, i, dequeueExpect.value, dequeuValue)
			}
		}

		length := aCase.queueInstance.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected value %v, got %v", caseNumber, aCase.expectedLength, length)
		}
	}

	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("Dequeue on a nil queue should cause a panic, did not")
		}
	}()
	nilQueue.Dequeue()
}

//*************** Concurrency Test ***************

//Test concurrent access to the Queue. Run with `go test -race` for better race detection.
func TestConcurrency(t *testing.T) {
	aQueue := NewQueue()
	aQueue.Enqueue(0)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go bombardQueue(aQueue, &wg)
	}

	wg.Wait()

	//Check there is exactly one value remaining in the queue
	lastRemainingValue, err := aQueue.Dequeue()
	if err != nil || lastRemainingValue != "-" {
		t.Errorf("Queue returning incorrect value after concurrent access")
	}
	_, err = aQueue.Dequeue()
	if err == nil {
		t.Errorf("Queue should be empty, but doesn't return error")
	}
}

func bombardQueue(queue *Queue, wg *sync.WaitGroup) {
	queue.Length()
	queue.Enqueue("-")
	queue.Peek()
	queue.Dequeue()
	queue.Length()

	time.Sleep(time.Microsecond)
	wg.Done()
}
