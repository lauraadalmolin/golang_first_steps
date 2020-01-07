package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// slices are basically dinamicaly sized arrays
	var slice1 []int
	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := []string{}

	// it is also possible to declare slices with make
	slice4 := make([]int, 10, 20)
	// where 10 is the length of the slice and 20 is
	// the amount of storage available to the slice

	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 3: ", slice3)
	fmt.Println("Slice 4: ", slice4)

	// there is only one iterator in go, which is for
	// the for structure is really versatile

	for i := 0; i < 10; i++ {
		fmt.Printf("Loop %d\n", i)
	}

	j := 0
	for j < 10 {
		j++
	}
	fmt.Printf("Exited for with j = %d\n", j)

	// also, you can implement infinite loops:
	rand.Seed(time.Now().UnixNano())
	n := 0
	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}
	fmt.Printf("Exited infinite loop after: %d\n", n)

	// slicing slices
	slicedSlice := slice2[1:3]
	fmt.Println("Sliced slice", slicedSlice)

	slice2[1] = 9
	fmt.Println("Changing original slice...")
	fmt.Println("Original slice:", slice2)
	fmt.Println("Sliced slice:", slicedSlice)
}
