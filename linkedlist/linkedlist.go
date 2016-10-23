package linkedlist

//*************** Linked List Public Interface ***************

type LinkedList struct {
	zeroElement *element
	length      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{zeroElement: nil, length: 0}
}

func (ll *LinkedList) Length() int {
	return 0
}

//GetValue returns error when index is out of bound or LinkedList is nil
func (ll *LinkedList) GetValue(index int) (value interface{}, err error) {
	return nil, nil
}

//Append works on a nil list too.
func (ll *LinkedList) Append(newValue interface{}) {
	return
}

//Remove gives an error when index is out of bound.
func (ll *LinkedList) Remove(index int) (removedValue interface{}, err error) {
	return nil, nil
}

//InsertBefore returns an error when index is out of bound.
func (ll *LinkedList) InsertBefore(index int, newValue interface{}) (err error) {
	return nil
}

//InsertAfter returns an error when index is out of bound.
func (ll *LinkedList) InsertAfter(index int, newValue interface{}) (err error) {
	return nil
}

//*************** Element ***************

type element struct {
	value interface{}
	next  *element
}
