package linkedlist_test

import (
	. "datatypes/linkedlist"
	"fmt"
	"testing"
)

//Test variables
var (
	nilList        *LinkedList
	emptyList      *LinkedList
	oneElementList *LinkedList
	twoElementList *LinkedList
	tenElementList *LinkedList
	veryLongList   *LinkedList
)

const (
	LENGTH_OF_VERY_LONG_LIST = 10000
)

//Sets test variables to default values.
func setVariablesToDefaults() {
	nilList = nil
	emptyList = NewLinkedList()

	oneElementList = NewLinkedList()
	oneElementList.Append(0)

	twoElementList = NewLinkedList()
	twoElementList.Append(0)
	twoElementList.Append(1)

	tenElementList = NewLinkedList()
	for i := 0; i < 10; i++ {
		tenElementList.Append(fmt.Sprintf("%d", i))
	}

	veryLongList = NewLinkedList()
	for i := 0; i < LENGTH_OF_VERY_LONG_LIST; i++ {
		veryLongList.Append(i)
	}
}

func TestNewLinkedList(t *testing.T) {
	linkedL := NewLinkedList()
	if linkedL == nil {
		t.Fatalf("LinkedList initialization returns nil")
	}
}

func TestLength(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list           *LinkedList
		expectedLength int
	}{
		{nilList, 0},
		{emptyList, 0},
		{oneElementList, 1},
		{twoElementList, 2},
		{tenElementList, 10},
		{veryLongList, LENGTH_OF_VERY_LONG_LIST},
	}

	for i, aCase := range cases {
		length := aCase.list.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected %d, got %d", i, aCase.expectedLength, length)
		}
	}
}

func TestGetValue(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list          *LinkedList
		index         int
		expectedValue interface{}
		expectError   bool
	}{
		//Empty lists
		{nilList, 0, nil, true},
		{nilList, 3, nil, true},
		{emptyList, 0, nil, true},
		{emptyList, -1, nil, true},
		{emptyList, 0, nil, true},
		//Small lists
		{oneElementList, 0, 0, false},
		{oneElementList, 1, nil, true},
		{twoElementList, 0, 0, false},
		{twoElementList, 1, 1, false},
		{twoElementList, 2, nil, true},
		//Longer list
		{tenElementList, 0, "0", false},
		{tenElementList, 9, "9", false},
		{tenElementList, 10, nil, true},
		{tenElementList, 100, nil, true},
		{tenElementList, -1, nil, true},
		//Large list
		{veryLongList, 30, 30, false},
		{veryLongList, LENGTH_OF_VERY_LONG_LIST - 1, LENGTH_OF_VERY_LONG_LIST - 1, false},
		{veryLongList, LENGTH_OF_VERY_LONG_LIST, nil, true},
	}

	for i, aCase := range cases {
		value, err := aCase.list.GetValue(aCase.index)
		if value != aCase.expectedValue {
			t.Errorf("Error in case %d. Expected value %v, got %v", i, aCase.expectedValue, value)
		}

		if !aCase.expectError && err != nil {
			t.Errorf("Error in case %d. Expected no error, got %s", i, err.Error())
		}
		if aCase.expectError && err == nil {
			t.Errorf("Error in case %d. Expected an error, got no error", i)
		}
	}
}

func TestAppend(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list          *LinkedList
		appendedValue interface{}
		//Check list values at selected indexes
		expectedLength                  int
		expectedIndexValuePairs         map[int]interface{}
		expectedIndexErrorExpectedPairs map[int]bool
	}{
		//Check exactly element is succesfully appended to an empty LinkedList
		{emptyList, "0", 1, map[int]interface{}{0: "0", 1: nil, 2: nil}, map[int]bool{0: false, 1: true, 2: true}},
		//Append one more element to a LinkedList of length 1
		{oneElementList, 1, 2, map[int]interface{}{0: 0, 1: 1, 2: nil}, map[int]bool{0: false, 1: false, 2: true}},
		//Append one more value to the result of the previous case
		{oneElementList, 2, 3, map[int]interface{}{0: 0, 1: 1, 2: 2}, map[int]bool{0: false, 1: false, 2: false}},
		//Append a value to a huge list
		{veryLongList, -1, LENGTH_OF_VERY_LONG_LIST + 1, map[int]interface{}{0: 0, 10: 10, LENGTH_OF_VERY_LONG_LIST: -1}, map[int]bool{0: false, 10: false, LENGTH_OF_VERY_LONG_LIST: false}},
	}

	for caseNumber, aCase := range cases {
		if len(aCase.expectedIndexValuePairs) != len(aCase.expectedIndexErrorExpectedPairs) {
			t.Fatalf("Test setup error in case %d. Number of index-value pairs: %d, number of index-error pairs: %d", caseNumber, len(aCase.expectedIndexErrorExpectedPairs), len(aCase.expectedIndexValuePairs))
		}

		aCase.list.Append(aCase.appendedValue)

		for index, expectedValue := range aCase.expectedIndexValuePairs {
			//Check that we know whether to expect an error
			expectError, ok := aCase.expectedIndexErrorExpectedPairs[index]
			if !ok {
				t.Fatalf("Test setup error in case %d. Error expectation not defined for index %d", caseNumber, index)
			}

			value, err := aCase.list.GetValue(index)
			if value != expectedValue {
				t.Errorf("Error in case %d, index %d. Expected value %v, got %v", caseNumber, index, expectedValue, value)
			}

			if !expectError && err != nil {
				t.Errorf("Error in case %d, index %d. Expected no error, got %s", caseNumber, index, err.Error())
			}
			if expectError && err == nil {
				t.Errorf("Error in case %d, index %d. Expected an error, got no error", caseNumber, index)
			}

			length := aCase.list.Length()
			if length != aCase.expectedLength {
				t.Errorf("Error in case %d. Expected final length %d, got %d", aCase.expectedLength, length)
			}
		}
	}

	//Test that a can struct can be appended
	aStruct := struct{}{}
	tenElementList.Append(aStruct)
	value, err := tenElementList.GetValue(10)
	if err != nil {
		t.Error("Error extracting a struct from a linked list")
	}
	if _, ok := value.(struct{}); !ok {
		t.Errorf("Structure added or extracted from linked list incorrectly")
	}

	//Test insert into a nil list panics
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Inserting into a nil LinkedList should cause a panic, but didn't.")
		}
	}()
	nilList.Append("-")
}

type expectation struct {
	expectedValue interface{}
	expectError   bool
}

func TestRemove(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list                 *LinkedList
		removalIndexSequence []int
		expectedLength       int
		expectedReturns      []expectation
	}{
		//Removal from empty lists. Check there is an error and no panic
		{list: emptyList, removalIndexSequence: []int{0, 0}, expectedLength: 0, expectedReturns: []expectation{{expectedValue: nil, expectError: true}, {expectedValue: nil, expectError: true}}},
		{list: emptyList, removalIndexSequence: []int{-1}, expectedLength: 0, expectedReturns: []expectation{{expectedValue: nil, expectError: true}}},
		//Removal from non-empty lists
		{list: oneElementList, removalIndexSequence: []int{0, 1}, expectedLength: 0, expectedReturns: []expectation{{expectedValue: 0, expectError: false}, {expectedValue: nil, expectError: true}}},
		{list: twoElementList, removalIndexSequence: []int{1, 0}, expectedLength: 0, expectedReturns: []expectation{{expectedValue: 1, expectError: false}, {expectedValue: 0, expectError: false}}},
		{list: tenElementList, removalIndexSequence: []int{0}, expectedLength: 9, expectedReturns: []expectation{{expectedValue: "0", expectError: false}}},
		//Removals from the end of a very long list
		{list: veryLongList, removalIndexSequence: []int{LENGTH_OF_VERY_LONG_LIST - 1, LENGTH_OF_VERY_LONG_LIST - 1}, expectedLength: LENGTH_OF_VERY_LONG_LIST - 1, expectedReturns: []expectation{{expectedValue: LENGTH_OF_VERY_LONG_LIST - 1, expectError: false}, {expectedValue: nil, expectError: true}}},
	}

	for caseNumber, aCase := range cases {
		if len(aCase.removalIndexSequence) != len(aCase.expectedReturns) {
			t.Errorf("Test setup error in case %d. Removal sequence length %s, expected returns length %s", caseNumber, len(aCase.removalIndexSequence), len(aCase.expectedReturns))
		}

		for i, removeIndex := range aCase.removalIndexSequence {
			value, err := aCase.list.Remove(removeIndex)

			if err != nil && !aCase.expectedReturns[i].expectError {
				t.Errorf("Error in case %d. Expected no error, got error: %s", caseNumber, err.Error())
			}
			if err == nil && aCase.expectedReturns[i].expectError {
				t.Errorf("Error in case %d. Expected error, got no error", caseNumber)
			}
			if value != aCase.expectedReturns[i].expectedValue {
				t.Errorf("Error in case %d. Expected value %v, got %v", caseNumber, aCase.expectedReturns[i].expectedValue, value)
			}
		}
		length := aCase.list.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected length %d, got %d", caseNumber, aCase.expectedLength, length)
		}
	}

	//Test removal from a nil list panics
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Removing from a nil LinkedList should cause a panic, but didn't.")
		}
	}()
	nilList.Remove(0)
}

func TestInsertBefore(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list         *LinkedList
		index        int
		insertValue  interface{}
		expectError  bool
		expectLength int
		//Check correct values are in correct position after the insert (uses GetValue() method)
		expectAtIndexes map[int]expectation
	}{
		//Inserting into empty lists:
		//Invalid indexes
		{emptyList, 1, "won't go in", true, 0, map[int]expectation{0: expectation{expectedValue: nil, expectError: true}, 1: expectation{expectedValue: nil, expectError: true}}},
		{emptyList, -1, "won't go in", true, 0, map[int]expectation{0: expectation{expectedValue: nil, expectError: true}, 1: expectation{expectedValue: nil, expectError: true}}},
		//Valid indexes
		{emptyList, 0, 0, false, 1, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: nil, expectError: true}}},
		//Inserting into non-empty list:
		//Invalid indexes
		{oneElementList, 2, "won't go in", true, 1, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: nil, expectError: true}}},
		{veryLongList, LENGTH_OF_VERY_LONG_LIST + 1, "won't go in", true, LENGTH_OF_VERY_LONG_LIST, map[int]expectation{LENGTH_OF_VERY_LONG_LIST - 1: expectation{expectedValue: LENGTH_OF_VERY_LONG_LIST - 1, expectError: false}, LENGTH_OF_VERY_LONG_LIST: expectation{expectedValue: nil, expectError: true}}},
		//Valid indexes
		{oneElementList, 1, 1, false, 2, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: 1, expectError: false}, 2: expectation{expectedValue: nil, expectError: true}}},
		{oneElementList, 2, 2, false, 3, map[int]expectation{1: expectation{expectedValue: 1, expectError: false}, 2: expectation{expectedValue: 2, expectError: false}}},
		{twoElementList, 0, 0, false, 3, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: 0, expectError: false}}},
		{veryLongList, 10, "10", false, LENGTH_OF_VERY_LONG_LIST + 1, map[int]expectation{20: expectation{expectedValue: 19, expectError: false}, LENGTH_OF_VERY_LONG_LIST: expectation{expectedValue: LENGTH_OF_VERY_LONG_LIST - 1, expectError: false}, LENGTH_OF_VERY_LONG_LIST + 1: expectation{expectedValue: nil, expectError: true}}},
	}

	for caseNumber, aCase := range cases {
		err := aCase.list.InsertBefore(aCase.index, aCase.insertValue)

		if err != nil && !aCase.expectError {
			t.Errorf("Error in case %d. Expected no error, got error: %s", caseNumber, err.Error())
		}
		if err == nil && aCase.expectError {
			t.Errorf("Error in case %d. Expected error, got no error", caseNumber)
		}

		//Check that correct values are in correct positions after the insert
		for index, anExpectation := range aCase.expectAtIndexes {
			value, err := aCase.list.GetValue(index)

			if err != nil && !anExpectation.expectError {
				t.Errorf("Error in case %d, index %d. Expected no error, got error: %s", caseNumber, index, err.Error())
			}
			if err == nil && anExpectation.expectError {
				t.Errorf("Error in case %d, index %d. Expected error, got no error", caseNumber, index)
			}
			if value != anExpectation.expectedValue {
				t.Errorf("Error in case %d, index %d. Expected value %v, got %v", caseNumber, index, anExpectation.expectedValue, value)
			}
		}
		length := aCase.list.Length()
		if length != aCase.expectLength {
			t.Errorf("Error in case %d. Expected length %d, got %d", caseNumber, aCase.expectLength, length)
		}
	}

	//Test insert into a nil list panics
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("InsertBefore into a nil LinkedList should cause a panic, but didn't.")
		}
	}()
	nilList.Append("-")
}

func TestInsertAfter(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		list         *LinkedList
		index        int
		insertValue  interface{}
		expectError  bool
		expectLength int
		//Check correct values are in correct position after the insert (uses GetValue() method)
		expectAtIndexes map[int]expectation
	}{
		//Inserting into empty lists:
		//Invalid indexes

		{emptyList, 0, "won't go in", true, 0, map[int]expectation{0: expectation{expectedValue: nil, expectError: true}, 1: expectation{expectedValue: nil, expectError: true}}},
		{emptyList, 1, "won't go in", true, 0, map[int]expectation{0: expectation{expectedValue: nil, expectError: true}, 1: expectation{expectedValue: nil, expectError: true}}},
		//Valid indexes
		{emptyList, -1, 0, false, 1, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: nil, expectError: true}}},
		//Inserting into non-empty list:
		//Invalid indexes
		{oneElementList, 1, "won't go in", true, 1, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: nil, expectError: true}}},
		{veryLongList, LENGTH_OF_VERY_LONG_LIST, "won't go in", true, LENGTH_OF_VERY_LONG_LIST, map[int]expectation{LENGTH_OF_VERY_LONG_LIST - 1: expectation{expectedValue: LENGTH_OF_VERY_LONG_LIST - 1, expectError: false}, LENGTH_OF_VERY_LONG_LIST: expectation{expectedValue: nil, expectError: true}}},
		//Valid indexes
		{oneElementList, 0, 1, false, 2, map[int]expectation{0: expectation{expectedValue: 0, expectError: false}, 1: expectation{expectedValue: 1, expectError: false}, 2: expectation{expectedValue: nil, expectError: true}}},
		{veryLongList, 10, "11", false, LENGTH_OF_VERY_LONG_LIST + 1, map[int]expectation{20: expectation{expectedValue: 19, expectError: false}, LENGTH_OF_VERY_LONG_LIST: expectation{expectedValue: LENGTH_OF_VERY_LONG_LIST - 1, expectError: false}, LENGTH_OF_VERY_LONG_LIST + 1: expectation{expectedValue: nil, expectError: true}}},
	}

	for caseNumber, aCase := range cases {
		err := aCase.list.InsertAfter(aCase.index, aCase.insertValue)

		if err != nil && !aCase.expectError {
			t.Errorf("Error in case %d. Expected no error, got error: %s", caseNumber, err.Error())
		}
		if err == nil && aCase.expectError {
			t.Errorf("Error in case %d. Expected error, got no error", caseNumber)
		}

		//Check that correct values are in correct positions after the insert
		for index, anExpectation := range aCase.expectAtIndexes {
			value, err := aCase.list.GetValue(index)

			if err != nil && !anExpectation.expectError {
				t.Errorf("Error in case %d, index %d. Expected no error, got error: %s", caseNumber, index, err.Error())
			}
			if err == nil && anExpectation.expectError {
				t.Errorf("Error in case %d, index %d. Expected error, got no error", caseNumber, index)
			}
			if value != anExpectation.expectedValue {
				t.Errorf("Error in case %d, index %d. Expected value %v, got %v", caseNumber, index, anExpectation.expectedValue, value)
			}
		}
		length := aCase.list.Length()
		if length != aCase.expectLength {
			t.Errorf("Error in case %d. Expected length %d, got %d", caseNumber, aCase.expectLength, length)
		}
	}

	//Test insert into a nil list panics
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("InsertAfter into a nil LinkedList should cause a panic, but didn't.")
		}
	}()
	nilList.Append("-")
}

func Example() {
	aLinkedList := NewLinkedList()

	aLinkedList.Append(3)
	_ = aLinkedList.InsertBefore(0, "Hello!")
	aLinkedList.Append(3)
	_ = aLinkedList.InsertBefore(2, "Hello!")
	aLinkedList.InsertAfter(1, []int{1, 2, 3})
	value, err := aLinkedList.Remove(1)
	if err != nil {
		panic(err.Error())
		//Handle error...
	}
	numValue := value.(int)
	fmt.Printf("%d\n", numValue)

	//Output:	3
}
