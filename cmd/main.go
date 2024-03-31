package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"com.pavdevs.katatest/internal/calculator"
	"com.pavdevs.katatest/pkg/support"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the operation (e.g., 2 * 2 or V * II):")
	scanner.Scan()
	input := scanner.Text()

	scanParts := strings.Fields(input)

	if len(scanParts) != 3 {
		panic("Input does not match the expected format (e.g., 2 * 2 or V * II)")
	}

	a, b, operator := scanParts[0], scanParts[2], scanParts[1]

	isRomanSystem := support.IsRoman(a) && support.IsRoman(b)

	if support.IsComplexFormat(a, b) {
		panic("Input does not match the expected format")
	}

	var result int

	if isRomanSystem {
		result = calculator.Calculate(support.ToArabic(a), support.ToArabic(b), operator)
		fmt.Println(support.ToRoman(result))
	} else {

		numA, errA := strconv.Atoi(a)
		numB, errB := strconv.Atoi(b)

		if errA != nil || errB != nil {
			panic("Invalid input")
		}

		if (numA < 1 || numA > 10) || (numB < 1 || numB > 10) {
			panic("Input numbers must be between 1 and 10")
		}

		result = calculator.Calculate(numA, numB, operator)
		fmt.Println(result)
	}
}
