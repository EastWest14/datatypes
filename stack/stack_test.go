package stack_test

import (
	. "datatypes/stack"
	"fmt"
	"testing"
)

var (
	nilStack        *Stack
	emptyStack      *Stack
	oneElementStack *Stack
	tenElementStack *Stack
	veryLargeStack  *Stack
)

const LENGTH_OF_LARGE_STACK = 100000

//Sets test variables to default values.
func setVariablesToDefaults() {
	nilStack = nil

	emptyStack = NewStack()

	oneElementStack = NewStack()
	oneElementStack.Push(0)

	tenElementStack = NewStack()
	for i := 0; i < 10; i++ {
		tenElementStack.Push(fmt.Sprintf("%d", i))
	}

	veryLargeStack = NewStack()
	for j := 0; j < LENGTH_OF_LARGE_STACK; j++ {
		veryLargeStack.Push(j)
	}
}

//*************** Public Interface Test ***************

func TestNewStack(t *testing.T) {
	aStack := NewStack()
	if aStack == nil {
		t.Fatalf("Initialization of new stack fails")
	}
}

func TestLength(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		stackInstance  *Stack
		expectedLength int
	}{
		{stackInstance: nilStack, expectedLength: 0},
		{stackInstance: emptyStack, expectedLength: 0},
		{stackInstance: oneElementStack, expectedLength: 1},
		{stackInstance: tenElementStack, expectedLength: 10},
		{stackInstance: veryLargeStack, expectedLength: LENGTH_OF_LARGE_STACK},
	}

	for i, aCase := range cases {
		length := aCase.stackInstance.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected stack length %d, got %d", i, aCase.expectedLength, length)
		}
	}
}

func TestPeek(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		stackInstance *Stack
		expectedValue interface{}
		expectError   bool
	}{
		{stackInstance: nilStack, expectedValue: nil, expectError: true},
		{stackInstance: emptyStack, expectedValue: nil, expectError: true},
		{stackInstance: oneElementStack, expectedValue: 0, expectError: false},
		//Check peeked value stays consistent
		{stackInstance: oneElementStack, expectedValue: 0, expectError: false},
		{stackInstance: tenElementStack, expectedValue: 10, expectError: false},
		{stackInstance: veryLargeStack, expectedValue: LENGTH_OF_LARGE_STACK, expectError: false},
	}

	for i, aCase := range cases {
		peekValue, err := aCase.stackInstance.Peek()

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

/*func TestPeek(t *testing.T) {

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
}*/

func Example() {
	aStack := NewStack()

	aStack.Push("First value")
	aStack.Push("Second value")

	value, err := aStack.Pop()
	if err != nil {
		//Handle error...
	}
	length := aStack.Length()

	fmt.Printf("Popped value: %v, length: %d", value, length)
	//Output: Popped value: first value, length: 1
}
