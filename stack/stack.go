//Stack is an implementation of a LIFO stack.
//Provides methods to push, pop and lookup values.
//Can be used directly or wrapped inside a custom structure.
//Safe to use concurrently.
package stack

import (
	"errors"
	"sync"
)

//Make runtime asserts fatal
const (
	panic_on_internal_inconsistencies = true
)

//*************** Stack Public Interface ***************

//Stack is a LIFO stack. Goroutine safe.
type Stack struct {
	length     int
	topElement *element
	rwMutex    sync.RWMutex
}

//NewStack initializes an empty Stack. Recommended way of initialization.
func NewStack() *Stack {
	return &Stack{length: 0, topElement: nil}
}

//Length returns the current number of values in the stack. Returns 0 on an uninitialized stack.
func (s *Stack) Length() int {
	if s == nil {
		return 0
	}

	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	return s.lengthValue()
}

//Peek returns the value at the top of the stack without removing it.
//If the stack is empty or nil, returns an error.
func (s *Stack) Peek() (value interface{}, err error) {
	if s == nil {
		return nil, errors.New("Stack is nil")
	}

	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	length := s.lengthValue()
	//Empty stack case
	if length == 0 {
		if s.topElement != nil && panic_on_internal_inconsistencies {
			panic("Stack is suppose to be empty, but top element is not nil")
		}

		return nil, errors.New("Stack is empty")
	}

	topElement := s.topElement
	return topElement.value, nil
}

//Pop removes the value from the top of the stack. If the stack is empty, returns an error.
//Panics on an uninitialized stack.
func (s *Stack) Pop() (value interface{}, err error) {
	if s == nil {
		panic("Stack is nil")
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	length := s.lengthValue()
	//Empty stack case
	if length == 0 {
		if s.topElement != nil && panic_on_internal_inconsistencies {
			panic("Stack is suppose to be empty, but top element is not nil")
		}

		return nil, errors.New("Stack is empty")
	}

	//Replace top element
	topElement := s.topElement
	s.topElement = topElement.previousElement
	s.changeLength(-1)
	return topElement.value, nil
}

//Push ads value to the top of the stack.
//Panics on an uninitialized stack.
func (s *Stack) Push(value interface{}) {
	if s == nil {
		panic("Stack is nil")
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	newElement := newElement(value)
	//currentTop can be nil
	currentTop := s.topElement
	s.topElement = newElement
	newElement.previousElement = currentTop
	s.changeLength(1)
	return
}

//*************** Stack Internal Structure ***************

type element struct {
	value           interface{}
	previousElement *element
}

func newElement(value interface{}) *element {
	return &element{value: value}
}

//Internal lenght method with no locking.
func (s *Stack) lengthValue() (length int) {
	return s.length
}

func (s *Stack) changeLength(delta int) {
	s.length += delta

	if s.length < 0 && panic_on_internal_inconsistencies {
		panic("Stack has negative length")
	}
}
