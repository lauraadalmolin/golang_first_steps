package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	statistics := getStatistics(words)
	write(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		counter, found := statistics[firstLetter]
		if found {
			statistics[firstLetter] = counter + 1
		} else {
			statistics[firstLetter] = 1
		}
	}

	return statistics
}

func write(statistics map[string]int) {
	fmt.Println("Word counting for each following first letter:")

	for firstLetter, counter := range statistics {
		fmt.Printf("%s = %d\n", firstLetter, counter)
	}
}
