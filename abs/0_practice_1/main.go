package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	fmt.Println(Practice1(a, b, c, s))
}

func Practice1(a int, b int, c int, s string) string {
	return fmt.Sprintf("%d %s", a+b+c, s)
}
