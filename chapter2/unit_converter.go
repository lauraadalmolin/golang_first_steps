package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	originUnit := os.Args[len(os.Args)-1]
	originValues := os.Args[1 : len(os.Args)-1]

	var destinyUnit string

	if originUnit == "celsius" {
		destinyUnit = "fahrenheit"
	} else if originUnit == "kilometers" {
		destinyUnit = "miles"
	} else {
		fmt.Printf("%s is not supported", destinyUnit)
		os.Exit(1)
	}

	for i, v := range originValues {
		originValues, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"The value %s at position %d "+
					"is not a valid number\n", v, i)
			os.Exit(1)
		}

		var destinyValues float64

		if originUnit == "celsius" {
			destinyValues = originValues*1.8 + 32
		} else {
			destinyValues = originValues / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			originValues, originUnit,
			destinyValues, destinyUnit)
	}
}
