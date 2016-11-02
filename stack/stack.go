//Stack is an implementation of a LIFO stack.
//Provides methods to push, pop and lookup values.
//Can be used directly or wrapped inside a custom structure.
//Safe to use concurrently.
package stack

//Make runtime asserts fatal
const (
	panic_on_internal_inconsistencies = true
)

//*************** Queue Public Interface ***************

//Stack is a LIFO stack. Goroutine safe.
type Stack struct {
}

//NewStack initializes an empty Stack. Recommended way of initialization.
func NewStack() *Stack {
	return &Stack{}
}

//Length returns the current number of values in the stack. Returns 0 on an uninitialized stack.
func (s *Stack) Length() int {
	return 0
}

//Peek returns the value at the top of the stack without removing it.
//If the stack is empty or nil, returns an error.
func (s *Stack) Peek() (value interface{}, err error) {
	return nil, nil
}

//Pop removes the value from the top of the stack. If the stack is empty, returns an error.
//Panics on an uninitialized stack.
func (s *Stack) Pop() (value interface{}, err error) {
	return nil, nil
}

//Push ads value to the top of the stack.
//Panics on an uninitialized stack.
func (s *Stack) Push(value interface{}) {
	return
}
