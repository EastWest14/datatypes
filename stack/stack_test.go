package stack_test

import (
	. "datatypes/stack"
	"fmt"
)

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
