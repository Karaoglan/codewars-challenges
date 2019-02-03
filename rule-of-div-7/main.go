package main

import "fmt"

func main() {
	result := Seven(477557101)
	// []int{28, 7}
	fmt.Println(result)
}

func Seven(n int64) []int {
	steps := 0
	for (n / 100) > 0 {
		steps++
		unitsDigit := n % 10
		n = (n / 10) - unitsDigit * 2
	}

	return []int{int(n), steps}
}
