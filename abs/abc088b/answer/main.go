package main

import (
	"fmt"
	"sort"
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
	if len(numbers)%2 != 0 {
		numbers = append(numbers, 0)
	}
	sort.Ints(numbers)

	var answer int

	for i := 0; i < len(numbers)/2; i++ {
		answer += numbers[2*i+1] - numbers[2*i]
	}
	return answer
}
