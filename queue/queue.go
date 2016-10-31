//Queue is an implementation of a FIFO queue.
//Provides methods to enqueue, dequeue and lookup values.
//Can be used directly or wrapped inside a custom structure.
//Safe to use concurrently.
package queue

import (
	"errors"
	"sync"
)

//Make runtime asserts fatal
const (
	panic_on_internal_inconsistencies = true
)

//*************** Queue Public Interface ***************

//Queue is a FIFO queue. Goroutine safe.
type Queue struct {
	length          int
	frontOfTheQueue *element
	backOfTheQueue  *element
	rwMutex         sync.RWMutex
}

//NewQueue initializes an empty Queue. Recommended way of initialization.
func NewQueue() *Queue {
	return &Queue{}
}

//Length returns the current number of values in the queue. Returns 0 on an uninitialized Queue.
func (q *Queue) Length() int {
	if q == nil {
		return 0
	}

	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	//Call internal function to obtain the length value
	return q.lengthValue()
}

//Peek returns the value at the front of the queue without removing it.
//If the queue is empty or nil, returns an error.
func (q *Queue) Peek() (value interface{}, err error) {
	if q == nil {
		return nil, errors.New("Queue is nil")
	}

	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	length := q.lengthValue()
	//If queue is empty - Peek returns an eror
	if length == 0 {
		return nil, errors.New("Queue is empty")
	}

	//Get value of the front element and check for internal inconsistencies
	frontElement := q.frontOfTheQueue
	if frontElement == nil {
		if panic_on_internal_inconsistencies {
			panic("Front element is nil, suppose to be not nil")
		}
		return nil, errors.New("Front element is nil, suppose to be not nil")
	}

	return frontElement.value, nil
}

//Enqueue adds the value to back of the queue.
//Panics on an uninitialized queue.
func (q *Queue) Enqueue(value interface{}) {
	if q == nil {
		panic("Queue is nil")
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	newElem := newElement(value, nil)
	length := q.lengthValue()

	//If the length is zero - set the new element as front and back of the queue.
	if length == 0 {
		q.frontOfTheQueue = newElem
		q.backOfTheQueue = newElem
		q.changeLength(1)
		return
	}

	//Make current back element point to new element. Set new back element.
	newElem.previousElement = nil
	q.backOfTheQueue.previousElement = newElem
	q.backOfTheQueue = newElem
	q.changeLength(1)
}

//Dequeue removes the value from the front of the queue. If queue is empty, returns error.
//Panics on an uninitialized queue.
func (q *Queue) Dequeue() (valueRemoved interface{}, err error) {
	if q == nil {
		panic("Queue is nil")
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	length := q.lengthValue()

	//If queue is empty - return error
	if length == 0 {
		return nil, errors.New("Queue is already empty")
	}

	//If length is 1 - remember front element, remove front and back elements.
	if length == 1 {
		currentFrontElement := q.frontOfTheQueue
		if currentFrontElement == nil && panic_on_internal_inconsistencies {
			panic("The only element in the queue is nil")
		}

		q.frontOfTheQueue = nil
		q.backOfTheQueue = nil
		q.changeLength(-1)
		return currentFrontElement.value, nil
	}

	//If length is > 1 - remember front element, reset front as its previous element (could be nil).
	currentFrontElement := q.frontOfTheQueue
	//Check for an internal runtime inconsistency
	if currentFrontElement == nil && panic_on_internal_inconsistencies {
		panic("Front element of the queue is nil")
	}

	q.frontOfTheQueue = currentFrontElement.previousElement
	q.changeLength(-1)
	return currentFrontElement.value, nil
}

//*************** Queue Internal Structure ***************

type element struct {
	value           interface{}
	previousElement *element
}

func newElement(value interface{}, previousElement *element) *element {
	return &element{value: value, previousElement: previousElement}
}

//Internal lenght method with no locking.
func (q *Queue) lengthValue() (length int) {
	return q.length
}

func (q *Queue) changeLength(delta int) {
	q.length += delta

	if q.length < 0 && panic_on_internal_inconsistencies {
		panic("Queue has negative length")
	}
}
