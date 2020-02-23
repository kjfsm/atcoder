package main

import "fmt"

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
	count := 0
	return loop(numbers, count)
}

func loop(numbers []int, count int) int {
	newNumbers := make([]int, len(numbers))

	for i, number := range numbers {
		if number%2 != 0 {
			return count
		}
		newNumbers[i] = number / 2
	}

	return loop(newNumbers, count+1)
}
