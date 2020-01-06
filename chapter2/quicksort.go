// Receives a list of numbers
// Elects an element (x) and removes it from the list
// Breaks the list in two, one containing elements bigger than x and another containing elements smaller than x
// Orders both lists recursively
// Returns a list containing the smaller elements, x, and the bigger elements.package chapter2

// ##### Recursive software #####

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a valid number\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

// the idea of this function is to keep breaking the array in smaller and smaller amounts
// until it is completely ordered
func quicksort(numbers []int) []int {
	// since this is a recursive function this condition
	// is extremely necessary, to avoid infinite loops (aka stack overflow)
	if len(numbers) <= 1 {
		return numbers
	}

	newList := make([]int, len(numbers))
	// I believe copy(x, y) works almost like angular.merge(x,y)
	copy(newList, numbers)

	pivotIndex := len(newList) / 2
	pivot := newList[pivotIndex]
	newList = append(newList[:pivotIndex], newList[pivotIndex+1:]...)

	smaller, greater := chop(newList, pivot)

	return append(
		append(quicksort(smaller), pivot), quicksort(greater)...)
}

func chop(numbers []int, pivot int) (smaller []int, greater []int) {
	for _, n := range numbers {
		if n <= pivot {
			smaller = append(smaller, n)
		} else {
			greater = append(greater, n)
		}
	}
	return smaller, greater
}
