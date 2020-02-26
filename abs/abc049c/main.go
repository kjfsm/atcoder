package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if Calc(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func Calc(s string) bool {
	return true
}
