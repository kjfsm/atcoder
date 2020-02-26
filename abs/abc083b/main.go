package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(Calc(n, a, b))
}

func Calc(n, a, b int) int {
	var answer int

	for i := 1; i <= n; i++ {
		var sum int
		str := strconv.Itoa(i)
		for _, rune := range str {
			digit, _ := strconv.Atoi(string(rune))
			sum += digit
		}
		if sum >= a && sum <= b {
			answer += i
		}
	}

	return answer
}
