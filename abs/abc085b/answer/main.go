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
	sort.Ints(numbers)
	answer := 1
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			answer++
		}
	}
	return answer
}
