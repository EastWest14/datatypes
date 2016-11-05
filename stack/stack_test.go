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
		{stackInstance: tenElementStack, expectedValue: "9", expectError: false},
		{stackInstance: veryLargeStack, expectedValue: LENGTH_OF_LARGE_STACK - 1, expectError: false},
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

//Used to test return values of the Pop method
type popExpectation struct {
	value       interface{}
	expectError bool
}

//TestPush adds elements to the stack then removes some of them and checks the returned values and errors.
func TestPush(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		stackInstance           *Stack
		pushSequence            []interface{}
		expectedLengthAfterPush int
		popExpectation          []popExpectation
	}{
		//Add value to an empty stack
		{emptyStack, []interface{}{0}, 1, []popExpectation{{value: 0, expectError: false}, {value: nil, expectError: true}}},
		//Add two values to a one element stack
		{oneElementStack, []interface{}{1, 2}, 3, []popExpectation{{value: 2, expectError: false}, {value: 1, expectError: false}, {value: 0, expectError: false}, {value: nil, expectError: true}}},
		//Test re-filling a completely empty stack
		{oneElementStack, []interface{}{"New only element"}, 1, []popExpectation{{value: "New only element", expectError: false}, {value: nil, expectError: true}}},
		//Add one value to a very long stack
		{veryLargeStack, []interface{}{"last value"}, LENGTH_OF_LARGE_STACK + 1, []popExpectation{{value: "last value", expectError: false}}},
	}

	for caseNumber, aCase := range cases {
		//Push all the values
		for _, pushValue := range aCase.pushSequence {
			aCase.stackInstance.Push(pushValue)
		}

		length := aCase.stackInstance.Length()
		if length != aCase.expectedLengthAfterPush {
			t.Errorf("Error in case %d. Expected value %v, got %v", caseNumber, aCase.expectedLengthAfterPush, length)
		}

		for i, popExpect := range aCase.popExpectation {
			popValue, err := aCase.stackInstance.Pop()

			//Check error value is expected
			if !popExpect.expectError {
				if err != nil {
					t.Errorf("Error in case %d, pop %d. Expected no error, got %s", caseNumber, i, err.Error)
				}
			} else {
				if err == nil {
					t.Errorf("Error in case %d, pop %d. Expected an error, got no error", caseNumber, i)
				}
			}

			//Check popped value is correct
			if popValue != popExpect.value {
				t.Errorf("Error in case %d, pop %d. Expected value %v, got %v", caseNumber, i, popExpect.value, popValue)
			}
		}
	}

	//Test Push on a nil stack panics
	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("Push on a nil stack should cause a panic, did not")
		}
	}()
	nilStack.Push(0)
}

func TestPop(t *testing.T) {
	setVariablesToDefaults()
	cases := []struct {
		stackInstance   *Stack
		popExpectations []popExpectation
		expectedLength  int
	}{
		{emptyStack, []popExpectation{{value: nil, expectError: true}}, 0},
		//Try double pop on already empty stack
		{oneElementStack, []popExpectation{{value: 0, expectError: false}, {value: nil, expectError: true}, {value: nil, expectError: true}}, 0},
		{tenElementStack, []popExpectation{{value: "9", expectError: false}}, 9},
		{veryLargeStack, []popExpectation{{value: LENGTH_OF_LARGE_STACK - 1, expectError: false}}, LENGTH_OF_LARGE_STACK - 1},
	}

	for caseNumber, aCase := range cases {
		for i, popExpect := range aCase.popExpectations {
			popValue, err := aCase.stackInstance.Pop()

			//Check error value is expected
			if !popExpect.expectError {
				if err != nil {
					t.Errorf("Error in case %d, pop %d. Expected no error, got %s", caseNumber, i, err.Error)
				}
			} else {
				if err == nil {
					t.Errorf("Error in case %d, pop %d. Expected an error, got no error", caseNumber, i)
				}
			}

			//Check pop value is correct
			if popValue != popExpect.value {
				t.Errorf("Error in case %d, pop %d. Expected value %v, got %v", caseNumber, i, popExpect.value, popValue)
			}
		}

		//Check the final length of the Stack is correct
		length := aCase.stackInstance.Length()
		if length != aCase.expectedLength {
			t.Errorf("Error in case %d. Expected value %v, got %v", caseNumber, aCase.expectedLength, length)
		}
	}

	//Test Pop on a nil stack panics
	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("Pop on a nil stack should cause a panic, did not")
		}
	}()
	nilStack.Pop()
}

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
	//Output: Popped value: Second value, length: 1
}
