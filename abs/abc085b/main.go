package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(Calc(numbers))
}

func Calc(numbers []int) int {
	return 0
}
