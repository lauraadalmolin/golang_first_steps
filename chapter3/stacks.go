package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Stack created with size ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())

	stack.ToStack("Go")
	stack.ToStack(1234)
	stack.ToStack("Pedro")
	stack.ToStack(1.23)
	fmt.Println("Size after stacking four values: ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())

	for !stack.Empty() {
		v, _ := stack.Unstack()
		fmt.Println("Unstacking... ", v)
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Empty? ", stack.Empty())
	}

	_, err := stack.Unstack()
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) ToStack(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Unstack() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Empty Stack!")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil

}
