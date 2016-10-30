package queue

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
	//efer read unlock
	return q.lengthValue()
}

//Peek returns the value of top element without removing it from the queue.
//If the queue is empty or nil, returns an error.
func (q *Queue) Peek() (value interface{}, err error) {
	//Check the queue is not nil
	//If it is nil, return error

	//Issue read block
	//Defer write unlock

	//Find length
	//If the length is zero - return error and nil

	//Get the top element
	//If it is nil - internal inconsistency
	//Return value and nil for error
	return nil, nil
}

func (q *Queue) Enqueue(value interface{}) {
	//Check the queue is not nil
	//If it is nil, panic

	//Issue write block
	//Defer write unlock

	//Create a new element
	//Find length

	//If the length is zero - set the element as top and bottom,
	//Append length.

	//Make new element point to the previous of current top.
	//Set new top. Append length.
}

func (q *Queue) Dequeue() (valueRemoved interface{}, err error) {
	//Check the queue is not nil
	//If it is nil, panic

	//Issue write block
	//Defer write unlock

	//Find length

	//If length is 0 - return error

	//If length is 1 - remember top element, remove top and bottom.
	//Decrease length. Return element value and nil.

	//If length is > 1 - remember to element, reset top as its previous.
	//Decrease length. Return element value and nil.

	return nil, nil
}

//*************** Queue Internal Structure ***************

type element struct {
	value            interface{}
	preveiousElement *element
}

func newElement(value interface{}, previousElement *element) *element {
	return &element{value: value, preveiousElement: previousElement}
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
