package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	fmt.Println(Calc(a, b, c, x))
}

func Calc(a, b, c, x int) int {
	return 0
}
