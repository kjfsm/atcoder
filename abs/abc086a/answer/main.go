package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(Calc(a, b))
}

func Calc(a int, b int) string {
	if (a*b)%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
