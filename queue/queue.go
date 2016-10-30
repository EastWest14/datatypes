package queue

import (
	"errors"
)

//Make runtime asserts fatal
const (
	panic_on_internal_inconsistencies = true
)

//*************** Queue Public Interface ***************

type Queue struct {
	length           int
	topOfTheQueue    *element
	bottomOfTheQueue *element
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Length() int {
	if q == nil {
		return 0
	}

	//Read block
	//defer read unlock

	return q.lengthValue()
}

//Peek returns the value at the front without removing it from the queue.
//If the queue is empty or nil, returns an error.
func (q *Queue) Peek() (value interface{}, err error) {
	//Check the queue is not nil
	//If it is nil, return error
	if q == nil {
		return nil, errors.New("Queue is nil")
	}

	//Issue read block
	//Defer write unlock

	//Find length
	length := q.lengthValue()
	if length == 0 {
		return nil, errors.New("Queue is empty")
	}
	//If the length is zero - return error and nil

	//Get the top element
	//If it is nil - internal inconsistency
	//Return value and nil for error
	topElement := q.topOfTheQueue
	if topElement == nil {
		if panic_on_internal_inconsistencies {
			panic("Top element is nil, suppose to be not nil")
		}
		return nil, errors.New("Top element is nil, suppose to be not nil")
	}

	return topElement.value, nil
}

//Enqueue adds the value to back of the queue.
//Panics on an iuninitialized queue
func (q *Queue) Enqueue(value interface{}) {
	if q == nil {
		panic("Queue is nil")
	}

	//Issue write block
	//Defer write unlock

	//Create a new element
	newElem := newElement(value, nil)
	//Find length
	length := q.lengthValue()

	//If the length is zero - set the element as top and bottom,
	//Append length. Return.
	if length == 0 {
		q.topOfTheQueue = newElem
		q.bottomOfTheQueue = newElem
		q.changeLength(1)
		return
	}

	//Make current bottom point to new element.
	//Set new bottom. Append length.
	newElem.previousElement = nil
	q.bottomOfTheQueue.previousElement = newElem
	q.bottomOfTheQueue = newElem
	q.changeLength(1)
}

func (q *Queue) Dequeue() (valueRemoved interface{}, err error) {
	if q == nil {
		panic("Queue is nil")
	}

	//Issue write block
	//Defer write unlock

	length := q.lengthValue()

	//If length is 0 - return error
	if length == 0 {
		return nil, errors.New("Queue is already empty")
	}

	//If length is 1 - remember top element, remove top and bottom.
	//Decrease length. Return element value and nil.
	if length == 1 {
		currentTopElement := q.topOfTheQueue
		if currentTopElement == nil && panic_on_internal_inconsistencies {
			panic("The only element in the queue is nil")
		}

		q.topOfTheQueue = nil
		q.bottomOfTheQueue = nil
		q.changeLength(-1)
		return currentTopElement.value, nil
	}

	//If length is > 1 - remember to element, reset top as its previous.
	//Decrease length. Return element value and nil.
	currentTopElement := q.topOfTheQueue
	if currentTopElement == nil && panic_on_internal_inconsistencies {
		panic("Top element of the queue is nil")
	}

	q.topOfTheQueue = currentTopElement.previousElement
	q.changeLength(-1)
	return currentTopElement.value, nil
}

//*************** Queue Internal Structure ***************

type element struct {
	value           interface{}
	previousElement *element
}

func newElement(value interface{}, previousElement *element) *element {
	return &element{value: value, previousElement: previousElement}
}

func (q *Queue) lengthValue() (length int) {
	return q.length
}

func (q *Queue) changeLength(delta int) {
	q.length += delta

	if q.length < 0 && panic_on_internal_inconsistencies {
		panic("Queue has negative length")
	}
}
