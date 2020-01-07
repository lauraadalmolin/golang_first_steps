package main

import (
	"fmt"
)

func main() {
	// declaring arrays
	var array [3]int
	numbers := [5]int{1, 2, 3, 4, 5}
	numbers2 := [...]int{2, 3, 5, 7, 11}
	names := [2]string{}

	// when the values are not attributed, default empty values,
	// aka zero values are attributed - int: 0, string: "", bool: false
	fmt.Println(array)
	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(names)

	// to know the size of an array
	names[0] = "Laura"
	names[1] = "Dalmolin"
	fmt.Println("Size of array 'names': ", len(names))

	// multidimensional arrays
	var multi [2][2]int
	multi[0][0] = 1
	multi[0][1] = 0
	// alternative way
	multi[1][0], multi[1][1] = 0, 1

	fmt.Println("Multidimensional array: ", multi)

}
