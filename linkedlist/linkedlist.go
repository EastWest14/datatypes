package linkedlist

import (
	"errors"
)

const (
	PANIC_ON_INTERNAL_INCONSTENCIES = true
)

//*************** Linked List Public Interface ***************

type LinkedList struct {
	baseElement *element
	length      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{baseElement: nil, length: 0}
}

func (ll *LinkedList) Length() int {
	return ll.lengthValue()
}

//GetValue returns error when index is out of bound or LinkedList is nil
func (ll *LinkedList) GetValue(index int) (value interface{}, err error) {
	if ll == nil {
		return nil, errors.New("Linked list is nil")
	}
	if index < 0 {
		return nil, errors.New("Can't get element - index is negative")
	}
	if index >= ll.lengthValue() {
		return nil, errors.New("Can't get element - index is out of bound")
	}

	elem, err := ll.elementAtIndex(index)
	if err != nil {
		return nil, err
	}
	return elem.value, nil
}

//Append panics with nil input
func (ll *LinkedList) Append(newValue interface{}) {
	if ll == nil {
		panic("Trying to append a nil linked list")
	}
	length := ll.lengthValue()
	ll.insertElementBefore(length, newElement(newValue))
}

//Remove gives an error when index is out of bound.
func (ll *LinkedList) Remove(index int) (removedValue interface{}, err error) {
	if ll == nil {
		panic("Trying to remove from a nil linked list")
	}

	if index < 0 {
		return nil, errors.New("Can't remove element - index is negative")
	}
	length := ll.lengthValue()
	if index >= length {
		return nil, errors.New("Can't remove element - index is out of bound")
	}

	if index == 0 {
		if length == 1 {
			value := ll.baseElement.value
			ll.baseElement = nil
			ll.changeLength(-1)
			return value, nil
		}

		newBaseElement, err := ll.elementAtIndex(1)
		if err != nil {
			if PANIC_ON_INTERNAL_INCONSTENCIES {
				panic("Failed to get the new base element.")
			}
			return nil, errors.New("Failed to get the new base element")
		}
		value := ll.baseElement.value
		ll.baseElement = newBaseElement
		ll.changeLength(-1)
		return value, nil
	}

	elementRightBefore, err := ll.elementAtIndex(index - 1)
	if err != nil {
		if PANIC_ON_INTERNAL_INCONSTENCIES {
			panic("Failed to get element right before the one being removed")
		}
		return nil, err
	}
	value := elementRightBefore.next.value
	elementRightBefore.next = elementRightBefore.next.next
	ll.changeLength(-1)

	return value, nil
}

//InsertBefore returns an error when index is out of bound.
func (ll *LinkedList) InsertBefore(index int, newValue interface{}) error {
	if ll == nil {
		panic("Trying to append a nil linked list")
	}

	if index < 0 {
		return errors.New("Can't insert value before index - index is negative")
	}
	length := ll.lengthValue()
	if index > length {
		return errors.New("Can't insert value before index - index out of bound")
	}
	ll.insertElementBefore(index, newElement(newValue))
	return nil
}

//InsertAfter returns an error when index is out of bound.
func (ll *LinkedList) InsertAfter(index int, newValue interface{}) (err error) {
	if ll == nil {
		panic("Trying to append a nil linked list")
	}

	if index < -1 {
		return errors.New("Can't insert value after - index is below -1")
	}
	length := ll.lengthValue()
	if index >= length {
		return errors.New("Can't insert value after index - index out of bound")
	}
	ll.insertElementBefore(index+1, newElement(newValue))
	return nil
}

//*************** Internal Structure ***************

func (ll *LinkedList) lengthValue() int {
	if ll == nil {
		return 0
	}
	return ll.length
}

func (ll *LinkedList) changeLength(delta int) {
	ll.length += delta

	if PANIC_ON_INTERNAL_INCONSTENCIES {
		if ll.length < 0 {
			panic("Length of the linked list is negative.")
		}
	}
}

func (ll *LinkedList) elementAtIndex(index int) (elem *element, err error) {
	if index < 0 {
		if PANIC_ON_INTERNAL_INCONSTENCIES {
			panic("Trying to get element at a negative index.")
		}
		return nil, errors.New("Trying to get an element at a negative index")
	}

	if index == 0 {
		if ll.baseElement == nil {
			return nil, errors.New("Base element does not exist.")
		}
		return ll.baseElement, nil
	}
	currentElement := ll.baseElement
	for i := 0; i < index; i++ {
		if currentElement.next == nil {
			return nil, errors.New("Element outside list boundary")
		}
		currentElement = currentElement.next
	}
	return currentElement, nil
}

func (ll *LinkedList) insertElementBefore(index int, insertedElement *element) {
	if index == 0 {
		ll.setBaseElement(insertedElement)
		ll.changeLength(1)
		return
	}

	elementOneBefore, err := ll.elementAtIndex(index - 1)
	if err != nil {
		if PANIC_ON_INTERNAL_INCONSTENCIES {
			panic("Failed to access element that was suppose to be in bounds")
		}
		return
	}
	if elementOneBefore == nil && PANIC_ON_INTERNAL_INCONSTENCIES {
		panic("Returned element is nil")
	}
	insertedElement.setNextElement(elementOneBefore.next)
	elementOneBefore.setNextElement(insertedElement)
	ll.changeLength(1)
}

func (ll *LinkedList) setBaseElement(newBaseELement *element) {
	ll.baseElement = newBaseELement
}

type element struct {
	value interface{}
	next  *element
}

func newElement(value interface{}) *element {
	return &element{value: value}
}

func (el *element) setNextElement(nextElement *element) {
	el.next = nextElement
}
